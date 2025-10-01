package repository

import (
	"context"
	"failiverCheck/internal/app/dto"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

func (r *Repository) UploadComponentImg(ctx context.Context, dataImg dto.ComponentImgCreateDTO) (string, error) {
	var filePath string
	filePath, err := CreateNewFilePath(dataImg.FilePath)
	if err != nil {
		return "", err
	}

	putOptions := minio.PutObjectOptions{
		ContentType: dataImg.ContentType,
	}
	_, err = r.minio.Client.PutObject(ctx, r.minio.Config.Bucket, filePath, dataImg.File, dataImg.FileSize, putOptions)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("http://%v:%v/%v/%v", r.minio.Config.Host, r.minio.Config.Port, r.minio.Config.Bucket, filePath), nil
}

func (r *Repository) DeleteComponentImg(ctx context.Context, filePath string) error {
	err := r.minio.Client.RemoveObject(context.Background(), r.minio.Config.Bucket, filePath, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetUrlComponentImg(componentId uint) (string, error) {
	component, err := r.GetComponentById(int(componentId))
	if err != nil {
		return "", err
	}
	url := component.Img
	if url == "" {
		return "", nil
	}
	lenBucket := len(r.minio.Config.Bucket)
	indexBucket := strings.Index(url, r.minio.Config.Bucket)
	if indexBucket == -1 {
		return "", fmt.Errorf("not found bucket %v in url-path %v", r.minio.Config.Bucket, url)
	}
	filePath := url[indexBucket+lenBucket+1:]
	logrus.Info(filePath)
	return filePath, nil
}
