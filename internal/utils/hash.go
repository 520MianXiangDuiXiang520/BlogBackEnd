package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/satori/go.uuid"
)

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func UUID() string {
	return uuid.NewV4().String()
}
