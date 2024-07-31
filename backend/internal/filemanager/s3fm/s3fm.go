package s3fm

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type s3fm struct {
	region string
	bucket string

	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func New(region, bucket string) (*s3fm, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String(region),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("could not create session: %v", err)
	}

	fm := &s3fm{
		region:     region,
		bucket:     bucket,
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}

	return fm, nil
}

func (fm s3fm) Save(ctx context.Context, body io.Reader, filePath string) (string, error) {
	output, err := fm.uploader.UploadWithContext(context.Background(), &s3manager.UploadInput{
		Key:    aws.String(filePath),
		Bucket: aws.String(fm.bucket),
		Body:   body,
	})
	if err != nil {
		return "", fmt.Errorf("could not upload file to s3 bucket: %v", err)
	}

	return output.Location, nil
}

// Deletes a file at the given path.
func (fm s3fm) Delete(ctx context.Context, filePath string) error {
	return fmt.Errorf("not implemented")
}

// Gets a file at the given path.
func (fm s3fm) Get(ctx context.Context, filePath string) (io.Reader, error) {
	// r := io.Reader
	// fm.downloader.DownloadWithContext(ctx, )
	//
	return nil, fmt.Errorf("not implemented")
}
