package minio

import (
	"context"
	"fmt"
	"log"
	"os"
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
	info, err := m.client.FPutObject(ctx, AudioBucketName, id.String(), filePath,
		mc.PutObjectOptions{ContentType: "audio/wav"})
	if err != nil {
		return fmt.Errorf("minio.SaveAudio: %w", err)
	}
	fmt.Printf("info: %v\n", info)
	return nil
}

func (m *Minio) LoadAudio(ctx context.Context, id uuid.UUID) ([]byte, error) {
	fullPath := fmt.Sprintf("/tmp/%s.wav", id)
	err := m.client.FGetObject(ctx, AudioBucketName, id.String(), fullPath, mc.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio.LoadAudio: %w", err)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("minio.LoadAudio: %w", err)
	}

	return data, nil
}
