package utils

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println(Sha256("zjb"))
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
}
