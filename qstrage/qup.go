package qstrage

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const Bucketname = "learning-rie"

func GetClient(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return client, nil
}

func WriteFile(ctx context.Context, s3path string, content io.Reader) error {
	log.Printf("write S3 file: %s/%s", Bucketname, s3path)
	client, err := GetClient(ctx)
	if err != nil {
		return err
	}
	obj := &s3.PutObjectInput{
		Bucket: aws.String(Bucketname),
		Key:    aws.String(s3path),
		Body:   content,
	}
	if _, err := client.PutObject(ctx, obj); err != nil {
		return err
	}
	return nil
}

func ReadJson(filepath string) (string, error) {
	log.Println("Loasing json:", filepath)
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s:%v\n", filepath, err)
	}
	return string(bs), nil
}
