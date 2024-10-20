package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Context string

const S3ContextKey S3Context = "aws-s3"

func S3Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatal("Unable to load AWS SDK")
			}

			sdk := s3.NewFromConfig(config)
			ctx := context.WithValue(request.Context(), S3ContextKey, sdk)

			next.ServeHTTP(writer, request.WithContext(ctx))
		},
	)
}

func GetBucketContext(request *http.Request) (*s3.Client, error) {
	sdk, ok := request.Context().Value(S3ContextKey).(*s3.Client)
	if !ok {
		return nil, errors.New("Unable to retrieve S3 SDK")
	}

	return sdk, nil
}
