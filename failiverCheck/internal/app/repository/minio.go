package repository

import (
	"failiverCheck/internal/app/config"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type Minio struct {
	Client *minio.Client
	Bucket string
}

func NewMinio(config *config.Config, bucket string) (*Minio, error) {
	endpoint := fmt.Sprintf("%s:%d", config.Minio.Host, config.Minio.Port)
	fmt.Println(endpoint)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &Minio{Client: client, Bucket: bucket}, nil

}
