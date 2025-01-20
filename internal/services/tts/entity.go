package tts

import (
	"time"

	sherpa_onnx "github.com/k2-fsa/sherpa-onnx-go-linux"
)

type TtsModel struct {
	model *sherpa_onnx.OfflineTts
	timer *time.Timer
}
