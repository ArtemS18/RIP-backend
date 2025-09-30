package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) UploadComponentImg(ctx context.Context, file io.Reader, fileSize int64, contentType string) (minio.UploadInfo, error) {
	putOptions := minio.PutObjectOptions{
		ContentType: contentType,
	}
	uploadInfo, err := r.minio.Client.PutObject(ctx, r.minio.Bucket, "myobject", file, fileSize, putOptions)
	if err != nil {
		fmt.Println(err)
		return minio.UploadInfo{}, err
	}
	return uploadInfo, nil
}
