package tts

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/av-ugolkov/lingua-ai/internal/config"

	"github.com/google/uuid"
	sherpa "github.com/k2-fsa/sherpa-onnx-go-linux"
)

type (
	minio interface {
		UploadAudio(ctx context.Context, id uuid.UUID, filePath string) error
		LoadAudio(ctx context.Context, id uuid.UUID) ([]byte, error)
	}
)

type Service struct {
	tts   map[string]*sherpa.OfflineTts
	minio minio
	sid   int
}

// TODO init LLM only if is needed
func New(cfg config.Tts) *Service {
	tts := make(map[string]*sherpa.OfflineTts, len(cfg.Models))
	for k, v := range cfg.Models {
		tts[k] = sherpa.NewOfflineTts(&sherpa.OfflineTtsConfig{
			Model: sherpa.OfflineTtsModelConfig{
				Vits: sherpa.OfflineTtsVitsModelConfig{
					Model:       v.VitsModel,
					Lexicon:     v.VitsLexicon,
					Tokens:      v.VitsTokens,
					DataDir:     v.VitsDataDir,
					NoiseScale:  v.VitsNoiseScale,
					NoiseScaleW: v.VitsNoiseScaleW,
					LengthScale: v.VitsLengthScale,
				},
				NumThreads: int(cfg.NumThreads),
				Debug:      int(cfg.Debug),
				Provider:   cfg.Provider,
			},
		})
	}

	return &Service{
		tts: tts,
	}
}

func (s *Service) SetMinio(minio minio) *Service {
	s.minio = minio
	return s
}

func (s *Service) Close(_ context.Context) error {
	for _, v := range s.tts {
		sherpa.DeleteOfflineTts(v)
	}

	return nil
}

func (s *Service) GetAudio(ctx context.Context, id uuid.UUID, text string, lang string) ([]byte, error) {
	data, err := s.getAudioData(ctx, id)
	if err != nil {
		slog.Warn(fmt.Sprintf("tts.Service.GetAudio: %v", err))
	}
	if data != nil {
		return data, nil
	}

	tts, ok := s.tts[lang]
	if !ok {
		return nil, errors.New("tts.Service.GetAudio: language not found")
	}
	audio := tts.Generate(text, s.sid, 1.0)
	pathFile := fmt.Sprintf("/tmp/%s.wav", id)
	audio.Save(pathFile)

	err = s.minio.UploadAudio(ctx, id, pathFile)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.GetAudio: %w", err)
	}

	data, err = s.getAudioData(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.GetAudio: %w", err)
	}

	return data, nil
}

func (s *Service) getAudioData(ctx context.Context, id uuid.UUID) ([]byte, error) {
	data, err := s.minio.LoadAudio(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.getAudioData: %w", err)
	}
	return data, nil
}
