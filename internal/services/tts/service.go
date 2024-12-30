package tts

import (
	"github.com/av-ugolkov/lingua-ai/internal/config"
	sherpa_onnx "github.com/k2-fsa/sherpa-onnx-go-linux"
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
			Model: sherpa_onnx.OfflineTtsModelConfig{
				Vits: sherpa_onnx.OfflineTtsVitsModelConfig{
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

func (s *Service) Close() {
	for _, v := range s.tts {
		sherpa.DeleteOfflineTts(v)
	}
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
