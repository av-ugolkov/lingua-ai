package minio

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/av-ugolkov/lingua-ai/internal/config"

	"github.com/google/uuid"
	mc "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	AudioBucketName = "audio"
)

type Minio struct {
	client *mc.Client
}

func Init(cfg *config.Minio) *Minio {
	minioClient, err := mc.New(cfg.Addr(), &mc.Options{
		Creds:  credentials.NewStaticV4(cfg.RootUser, cfg.RootPsw, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	m := &Minio{
		client: minioClient,
	}

	m.CreateBucket(AudioBucketName)

	return m
}

func (m *Minio) CreateBucket(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := m.client.BucketExists(ctx, name)
	if err == nil && ok {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	err = m.client.MakeBucket(ctx, name, mc.MakeBucketOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Minio) UploadAudio(ctx context.Context, id uuid.UUID, filePath string) error {
	_, err := m.client.FPutObject(ctx, AudioBucketName, id.String(), filePath,
		mc.PutObjectOptions{ContentType: "audio/wav"})
	if err != nil {
		return fmt.Errorf("minio.SaveAudio: %w", err)
	}
	return nil
}

func (m *Minio) LoadAudio(ctx context.Context, id uuid.UUID) ([]byte, error) {
	obj, err := m.client.GetObject(ctx, AudioBucketName, id.String(), mc.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio.LoadAudio: %w", err)
	}

	dataStat, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("minio.LoadAudio: %w", err)
	}

	data := make([]byte, dataStat.Size)
	_, err = obj.Read(data)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("minio.LoadAudio: %w", err)
	}

	return data, nil
}
