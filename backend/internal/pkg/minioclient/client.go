package minioclient

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"time"
)

const (
	healthCheckDuration = time.Second * 30
)

type Client struct {
	cli *minio.Client
}

func New(ctx context.Context, hostPort, idKey, secretKey string, useSSL bool) (*Client, error) {
	minioClient, err := minio.New(hostPort, &minio.Options{
		Creds:        credentials.NewStaticV4(idKey, secretKey, ""),
		Secure:       useSSL,
		CustomMD5:    nil,
		CustomSHA256: nil,
	})
	if err != nil {
		return nil, err
	}

	_, err = minioClient.HealthCheck(healthCheckDuration)
	if err != nil {
		return nil, err
	}

	if err = initBuckets(ctx, minioClient); err != nil {
		return nil, err
	}

	return &Client{cli: minioClient}, nil
}

func (c *Client) PutObject(
	ctx context.Context,
	reader io.Reader,
	size int64,
	bucketName, contentType string,
) (_ string, err error) {
	objectId := generateObjectName()

	opts := minio.PutObjectOptions{ContentType: contentType}

	_, err = c.cli.PutObject(ctx, bucketName, objectId, reader, size, opts)
	if err != nil {
		return "", fmt.Errorf("failed put file %s: %v", objectId, err)
	}

	return objectId, nil
}

func (c *Client) GetObject(ctx context.Context, objectId, bucketName string) (io.Reader, string, error) {
	reader, err := c.cli.GetObject(ctx, bucketName, objectId, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", fmt.Errorf("failed get object %s: %v", objectId, err)
	}

	stat, err := reader.Stat()
	if err != nil {
		return nil, "", fmt.Errorf("failed get object stat %s: %v", objectId, err)
	}

	return reader, stat.ContentType, nil
}

func generateObjectName() string {
	return uuid.New().String()
}

func initBuckets(ctx context.Context, cli *minio.Client) error {
	buckets := []string{
		model.ImageBucketName,
	}

	for _, bucket := range buckets {
		exists, err := cli.BucketExists(ctx, bucket)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		err = cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil
		}
	}

	return nil
}
