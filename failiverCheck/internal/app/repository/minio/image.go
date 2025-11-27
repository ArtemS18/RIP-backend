package minio

import (
	"context"
	"failiverCheck/internal/app/dto"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

func (m *Minio) UploadComponentImg(ctx context.Context, dataImg dto.ComponentImgCreateDTO) (string, error) {
	var filePath string
	filePath, err := CreateNewFilePath(dataImg.FilePath)
	if err != nil {
		return "", err
	}

	putOptions := minio.PutObjectOptions{
		ContentType: dataImg.ContentType,
	}
	_, err = m.Client.PutObject(ctx, m.Config.Bucket, filePath, dataImg.File, dataImg.FileSize, putOptions)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("http://%v:%v/%v/%v", m.Config.Host, m.Config.Port, m.Config.Bucket, filePath), nil
}

func (m *Minio) DeleteComponentImg(ctx context.Context, imgUrl *string) error {
	filePath, err := m.GetUrlComponentImg(imgUrl)
	if err != nil {
		return err
	}
	err = m.Client.RemoveObject(context.Background(), m.Config.Bucket, filePath, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (m *Minio) GetUrlComponentImg(imgUrl *string) (string, error) {
	url := *imgUrl
	if url == "" {
		return "", nil
	}
	lenBucket := len(m.Config.Bucket)
	indexBucket := strings.Index(url, m.Config.Bucket)
	if indexBucket == -1 {
		return "", fmt.Errorf("not found bucket %v in url-path %v", m.Config.Bucket, url)
	}
	filePath := url[indexBucket+lenBucket+1:]
	logrus.Info(filePath)
	return filePath, nil
}
