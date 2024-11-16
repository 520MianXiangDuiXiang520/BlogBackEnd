package logfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"slices"
	"testing"
)

func TestNewLogFile(t *testing.T) {
	fileName := "./test.log"
	f, err := NewLogFile(fileName)
	assert.Nil(t, err)
	data := []byte("test")
	_, err = f.Write(data)
	assert.Nil(t, err)
	readData, err := os.ReadFile(fileName)
	assert.Nil(t, err)
	assert.Equal(t, readData, data)
	err = os.Remove(fileName)
	assert.Nil(t, err)
}

func TestNewLogFile_Split(t *testing.T) {
	err := os.Mkdir("test", 0750)
	assert.Nil(t, err)

	defer func() {
		_ = os.RemoveAll("test")
	}()

	fileName := "./test/test.log"
	f, err := NewLogFile(fileName, WithSplit(4))
	assert.Nil(t, err)

	data := []byte("test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	data = []byte("test test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	data = []byte("tes")
	_, err = f.Write(data)
	assert.Nil(t, err)
	data = []byte("test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	data = []byte("111")
	_, err = f.Write(data)
	assert.Nil(t, err)

	lst, err := os.ReadDir("./test")
	assert.Nil(t, err)

	slices.SortFunc(lst, func(a, b os.DirEntry) int {
		infoA, _ := a.Info()
		infoB, _ := b.Info()
		fmt.Println(infoA.ModTime(), infoB.ModTime(), infoA.Name(), infoB.Name())
		return infoA.ModTime().Nanosecond() - infoB.ModTime().Nanosecond()
	})

	expected := []string{"test", "test test", "testest", "111"}
	for i, entry := range lst {
		data, err := os.ReadFile("test/" + entry.Name())
		assert.Nil(t, err)
		assert.Equalf(t, expected[i], string(data), "%s", entry.Name())
	}
}
