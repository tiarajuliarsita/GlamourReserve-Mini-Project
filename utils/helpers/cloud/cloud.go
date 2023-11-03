package cloud

import (
	"context"
	"fmt"
	cfg "glamour_reserve/app/config"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func ConfigCloud() *s3.Client {

	accountId := cfg.CloudAccount()
	accessKeyId := cfg.CloudKeyId()
	accessKeySecret := cfg.CloudKeySecret()

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}

func UploadFile(file *multipart.FileHeader, client *s3.Client) string {

	bucketName := cfg.CloudBucket()

	src, err := file.Open()
	randomKey := uuid.New().String()
	input := &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &randomKey,
		Body:        src,
		ContentType: aws.String("image/png"),
	}
	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	publiccostume := fmt.Sprintf("https://storage.glamourreserve.my.id/%s", randomKey)
	return publiccostume

}
