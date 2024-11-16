package logfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"slices"
	"testing"
	"time"
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
	time.Sleep(time.Nanosecond * 10)

	data := []byte("test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	time.Sleep(time.Nanosecond * 10)
	data = []byte("test test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	time.Sleep(time.Nanosecond * 10)
	data = []byte("tes")
	_, err = f.Write(data)
	assert.Nil(t, err)
	data = []byte("test")
	_, err = f.Write(data)
	assert.Nil(t, err)

	time.Sleep(time.Nanosecond * 10)
	data = []byte("111")
	_, err = f.Write(data)
	assert.Nil(t, err)

	lst, err := os.ReadDir("./test")
	assert.Nil(t, err)

	slices.SortFunc(lst, func(a, b os.DirEntry) int {
		infoA, err := a.Info()
		assert.Nil(t, err)
		infoB, err := b.Info()
		assert.Nil(t, err)
		fmt.Println(infoA.ModTime(), infoB.ModTime(), infoA.Name(), infoB.Name())
		if infoA.ModTime().Before(infoB.ModTime()) {
			return -1
		} else if infoA.ModTime().After(infoB.ModTime()) {
			return 1
		} else {
			return 0
		}
	})

	expected := []string{"test", "test test", "testest", "111"}
	for i, entry := range lst {
		data, err := os.ReadFile("test/" + entry.Name())
		assert.Nil(t, err)
		assert.Equalf(t, expected[i], string(data), "%s", entry.Name())
	}
}
