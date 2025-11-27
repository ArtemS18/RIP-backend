package minio

import (
	"failiverCheck/internal/app/config"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type Minio struct {
	Client *minio.Client
	Config *config.MinioConfig
}

func NewMinio(config *config.MinioConfig) (*Minio, error) {
	endpoint := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Println(endpoint)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &Minio{Client: client, Config: config}, nil

}
