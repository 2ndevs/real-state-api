package application

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"main/core"
	"main/domain/entities"
	"main/utils"
	"mime/multipart"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdatePropertyService struct {
	Validated *validator.Validate
	Database  *gorm.DB
	Bucket    *s3.Client
}

type UpdatePropertyRequest struct {
	Property       entities.Property
	DeleteIds      *[]string
	UploadedImages []*multipart.FileHeader
}

type UploadImageRequest struct {
	wg     *sync.WaitGroup
	image  *multipart.FileHeader
	bucket *s3.Client
}

type DeleteImageRequest struct {
	wg     *sync.WaitGroup
	image  string
	bucket *s3.Client
}

type UploadChannelResponse struct {
	value *string
	err   error
}

func (self *UpdatePropertyService) Execute(payload UpdatePropertyRequest, propertyID uint64) (*entities.Property, error) {
	property := payload.Property

	validationErr := self.Validated.Struct(property)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingProperty *entities.Property
	query := self.Database.Model(&entities.Property{}).Where("id = ? and deleted_at IS NULL", propertyID)

	existingPropertyDatabaseResponse := query.First(&existingProperty)
	if errors.Is(existingPropertyDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingPropertyDatabaseResponse.Error != nil {
		return nil, existingPropertyDatabaseResponse.Error
	}

	list := utils.WatchedList[string]{}
	list.Create(existingProperty.PreviewImages)

	for _, item := range *payload.DeleteIds {
		list.Remove(item)
	}

	var waitUpload sync.WaitGroup
	waitUpload.Add(len(payload.UploadedImages))

	for _, image := range payload.UploadedImages {
		channel := make(chan UploadChannelResponse, 1)
		uploadImageRequest := UploadImageRequest{
			wg:     &waitUpload,
			bucket: self.Bucket,
			image:  image,
		}

		go handleUpload(uploadImageRequest, channel)

		channelResponse := <-channel
		if channelResponse.err != nil {
			return nil, channelResponse.err
		}

		if channelResponse.value != nil {
			list.Add(*channelResponse.value)
		}
	}

	waitUpload.Wait()

	var waitDeletion sync.WaitGroup
	waitDeletion.Add(len(list.GetRemoved()))

	for _, image := range list.GetRemoved() {
		deletionChannel := make(chan error, 1)
		deleteImageRequest := DeleteImageRequest{
			wg:     &waitDeletion,
			image:  image,
			bucket: self.Bucket,
		}

		go handleDelete(deleteImageRequest, deletionChannel)

		err := <-deletionChannel
		if err != nil {
			return nil, err
		}
	}

	waitDeletion.Wait()

	property.ID = existingProperty.ID
	property.PreviewImages = list.GetItems()

	updatePropertyTransaction := self.Database.Save(&property)
	if updatePropertyTransaction.Error != nil {
		return nil, updatePropertyTransaction.Error
	}

	return &property, nil
}

func handleDelete(params DeleteImageRequest, channel chan<- error) {
	defer params.wg.Done()

	if len(params.image) == 0 || params.image == "fallback.jpg" {
		channel <- nil
		return
	}

	item := s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(params.image),
	}

	_, err := params.bucket.DeleteObject(context.TODO(), &item)
	if err != nil {
		channel <- err
		return
	}

	channel <- nil
}

func handleUpload(params UploadImageRequest, channel chan<- UploadChannelResponse) {
	response := UploadChannelResponse{
		value: nil,
		err:   nil,
	}

	defer params.wg.Done()

	if params.image == nil {
		channel <- response
		return
	}

	content, err := params.image.Open()
	defer content.Close()

	if err != nil {
		response.err = err
		channel <- response

		return
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, content)
	if err != nil {
		response.err = err
		channel <- response

		return
	}

	ext := strings.Split(params.image.Filename, ".")[1]
	if len(ext) == 0 {
		response.err = core.ImageUploadError
		channel <- response

		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		response.err = err
		channel <- response

		return
	}

	name := id.String()
	if len(name) == 0 {
		response.err = core.ImageUploadError
		channel <- response

		return
	}

	fileName := fmt.Sprintf("%v.%v", name, ext)

	item := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer.Bytes()),
	}

	_, err = params.bucket.PutObject(context.TODO(), item)
	if err != nil {
		response.err = err
		channel <- response

		return
	}

	response.value = &fileName
	channel <- response
}
