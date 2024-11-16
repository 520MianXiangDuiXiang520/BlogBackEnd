package logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	out := bytes.NewBufferString("")
	log := NewLogger(out,
		WithLevel(slog.LevelDebug),
		//WithSource(),
		//WithTime(),
		//WithTimeFormat(time.DateTime),
	)

	log.Info("Success", "a", 1, slog.Int("po", 100))
	assert.Equal(t, "level=INFO msg=Success a=1 po=100\n", out.String())
}
