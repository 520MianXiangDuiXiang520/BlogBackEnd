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

	infos := make([]os.FileInfo, 0)
	for _, l := range lst {
		info, err := l.Info()
		assert.Nil(t, err)
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		// 文件没有关闭 在某些系统上拿到的最后修改时间不准确
		if infos[i].Name() == "test.log" {
			return false
		}
		if infos[j].Name() == "test.log" {
			return true
		}
		return infos[i].ModTime().Before(infos[j].ModTime())
	})

	for _, entry := range infos {
		fmt.Printf(" %s %s\n", entry.ModTime(), entry.Name())
	}

	expected := []string{"test", "test test", "testest", "111"}
	for i, entry := range infos {
		data, err := os.ReadFile("test/" + entry.Name())
		assert.Nil(t, err)
		assert.Equalf(t, expected[i], string(data), "%s", entry.Name())
	}
}
