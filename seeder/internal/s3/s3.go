package s3

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/pkorobeinikov/seeder/seeder"
)

const (
	SeederS3EndpointEnv        = "SEEDER_S3_ENDPOINT"
	SeederS3AccessKeyIDEnv     = "SEEDER_S3_ACCESS_KEY_ID"
	SeederS3SecretAccessKeyEnv = "SEEDER_S3_SECRET_ACCESS_KEY"

	defaultLocation = "us-east-1"
)

func Seed(ctx context.Context, cfg seeder.Config) error {

	useSSL := false

	s3Endpoint, found := os.LookupEnv(SeederS3EndpointEnv)
	if !found {
		return errors.New("s3 access id key not set")
	}

	s3AccessKeyID, found := os.LookupEnv(SeederS3AccessKeyIDEnv)
	if !found {
		return errors.New("s3 access id key not set")
	}

	s3SecretAccessKey, found := os.LookupEnv(SeederS3SecretAccessKeyEnv)
	if !found {
		return errors.New("s3 secret access key not set")
	}

	minioClient, err := minio.New(s3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKeyID, s3SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return errors.Wrap(err, "minio new client")
	}

	err = minioClient.MakeBucket(
		ctx,
		cfg.Bucket,
		minio.MakeBucketOptions{Region: defaultLocation},
	)
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, cfg.Bucket)
		if errBucketExists == nil && !exists {
			return errors.Wrap(err, "minio: bucket not exists")
		}
	}

	_, err = minioClient.FPutObject(ctx, cfg.Bucket, cfg.ObjectName, cfg.File, minio.PutObjectOptions{
		ContentType:     cfg.Option.ContentType,
		ContentEncoding: cfg.Option.ContentEncoding,
	})
	if err != nil {
		return errors.Wrap(err, "minio: put object")
	}

	return nil
}

func init() {
	seeder.DefaultRegistry().RegisterSeeder(Seed, "s3")
}
