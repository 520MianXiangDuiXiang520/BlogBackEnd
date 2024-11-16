package logfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
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

	sort.Slice(lst, func(i, j int) bool {
		infoA, err := lst[i].Info()
		assert.Nil(t, err)
		infoB, err := lst[j].Info()
		assert.Nil(t, err)
		return infoA.ModTime().Before(infoB.ModTime())
	})

	for _, entry := range lst {
		fmt.Printf("%s\n", entry.Name())
	}

	expected := []string{"test", "test test", "testest", "111"}
	for i, entry := range lst {
		data, err := os.ReadFile("test/" + entry.Name())
		assert.Nil(t, err)
		assert.Equalf(t, expected[i], string(data), "%s", entry.Name())
	}
}
