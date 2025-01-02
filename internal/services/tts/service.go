package tts

import (
	"context"

	"github.com/av-ugolkov/lingua-ai/internal/config"

	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

type Service struct {
	tts map[string]*sherpa.OfflineTts
	sid int
}

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

func (s *Service) Close(_ context.Context) error {
	for _, v := range s.tts {
		sherpa.DeleteOfflineTts(v)
	}

	return nil
}

func (s *Service) GetAudio(text string, lang string) []float32 {
	tts, ok := s.tts[lang]
	if !ok {
		return nil
	}
	audio := tts.Generate(text, s.sid, 1.0)
	audio.Save("audio.wav")
	return audio.Samples
}
