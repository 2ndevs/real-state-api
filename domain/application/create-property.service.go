package application

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"main/core"
	"main/domain/entities"
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

type CreatePropertyService struct {
	Validated *validator.Validate
	Database  *gorm.DB
	Bucket    *s3.Client
}

type CreatePropertyServiceRequest struct {
	entities.Property
	PreviewImages []*multipart.FileHeader `validated:"required,min=1"`
}

type ImageUploadChannelResponse struct {
	value *string
	err   error
}

func (self *CreatePropertyService) Execute(property CreatePropertyServiceRequest) (*entities.Property, error) {
	validationErr := self.Validated.Struct(property)
	if validationErr != nil {
		log.Println(validationErr.Error())
		return nil, core.InvalidParametersError
	}

	var wait sync.WaitGroup
	var images []string
	wait.Add(len(property.PreviewImages))

	for _, image := range property.PreviewImages {
		channel := make(chan ImageUploadChannelResponse, 1)

		go func(image *multipart.FileHeader, channel chan<- ImageUploadChannelResponse) {
			response := ImageUploadChannelResponse{
				value: nil,
				err:   nil,
			}

			content, err := image.Open()

			defer content.Close()
			defer wait.Done()

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

			ext := strings.Split(image.Filename, ".")[1]
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
				response.err = err
				channel <- response
				return
			}

			fileName := fmt.Sprintf("%v.%v", name, ext)

			item := &s3.PutObjectInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
				Key:    aws.String(fileName),
				Body:   bytes.NewReader(buffer.Bytes()),
			}

			_, err = self.Bucket.PutObject(context.TODO(), item)
			if err != nil {
				response.err = core.ImageUploadError
				channel <- response
				return
			}

			response.value = &fileName
			channel <- response
		}(image, channel)

		response := <-channel

		if response.err != nil {
			return nil, response.err
		}

		images = append(images, *response.value)
	}

	wait.Wait()

	property.Property.PreviewImages = append(property.Property.PreviewImages, images...)

	if len(property.Property.PreviewImages) == 0 {
		property.Property.PreviewImages = []string{"fallback.jpg"}
	}

	createPropertyTransaction := self.Database.Create(&property.Property)
	if createPropertyTransaction.Error != nil {
		return nil, createPropertyTransaction.Error
	}

	return &property.Property, nil
}
