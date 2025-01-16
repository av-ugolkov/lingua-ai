package tts

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/av-ugolkov/lingua-ai/internal/config"

	"github.com/google/uuid"
	sherpa "github.com/k2-fsa/sherpa-onnx-go-linux"
)

const (
	speed = 1.0
)

type (
	minio interface {
		UploadAudio(ctx context.Context, id uuid.UUID, filePath string) error
		LoadAudio(ctx context.Context, id uuid.UUID) ([]byte, error)
	}
)

type TtsModel struct {
	model *sherpa.OfflineTts
	timer *time.Timer
}

type Service struct {
	tts     map[string]*TtsModel
	models  map[string]*sherpa.OfflineTtsConfig
	minio   minio
	sid     int
	timeout time.Duration

	mx sync.RWMutex
}

func New(cfg config.Tts) *Service {
	models := make(map[string]*sherpa.OfflineTtsConfig, len(cfg.Models))

	for k, v := range cfg.Models {
		models[k] = &sherpa.OfflineTtsConfig{
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
		}
	}

	return &Service{
		tts:     make(map[string]*TtsModel, len(cfg.Models)),
		models:  models,
		timeout: cfg.Timeout,
	}
}

func (s *Service) SetMinio(minio minio) *Service {
	s.minio = minio
	return s
}

func (s *Service) Close(_ context.Context) error {
	for _, v := range s.tts {
		sherpa.DeleteOfflineTts(v.model)
	}

	return nil
}

func (s *Service) getModel(lang string) (*sherpa.OfflineTts, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	m, ok := s.tts[lang]
	if ok {
		s.tts[lang].timer.Reset(s.timeout)
		return m.model, nil
	}

	s.tts[lang] = &TtsModel{
		model: sherpa.NewOfflineTts(s.models[lang]),
		timer: time.AfterFunc(s.timeout, func() {
			s.mx.Lock()
			defer s.mx.Unlock()
			sherpa.DeleteOfflineTts(s.tts[lang].model)
			delete(s.tts, lang)
		}),
	}
	return s.tts[lang].model, nil
}

func (s *Service) GetAudio(ctx context.Context, id uuid.UUID, text string, lang string) ([]byte, error) {
	data, err := s.getAudioData(ctx, id)
	if err != nil {
		slog.Warn(fmt.Sprintf("tts.Service.GetAudio: %v", err))
	}
	if data != nil {
		return data, nil
	}

	tts, err := s.getModel(lang)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.GetAudio: %v", err)
	}

	if !strings.HasPrefix(text, " ") {
		text = fmt.Sprintf(" %s", text)
	}
	audio := tts.Generate(text, s.sid, speed)

	data, err = s.uploadAudioData(ctx, id, audio)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.GetAudio: %w", err)
	}

	return data, nil
}

func (s *Service) uploadAudioData(ctx context.Context, id uuid.UUID, audio *sherpa.GeneratedAudio) ([]byte, error) {
	pathFile := fmt.Sprintf("/tmp/%s.wav", id)

	if ok := audio.Save(pathFile); !ok {
		return nil, fmt.Errorf("tts.Service.uploadAudioData: file not saved")
	}

	defer os.Remove(pathFile)

	err := s.minio.UploadAudio(ctx, id, pathFile)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.uploadAudioData: %w", err)
	}

	data, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, fmt.Errorf("tts.Service.uploadAudioData: %w", err)
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
