package config

import (
	"JuneBlog/patch/logger"
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	*Common
	Db DbConfig `json:"db"`
}

func NewConfig() *Config {
	return &Config{Common: NewCommon()}
}

var (
	G    *Config
	once sync.Once
)

func InitConfig(path string) error {
	once.Do(func() {
		data, err := os.ReadFile(path)
		if err != nil {
			logger.Panic("read config fail", "path", path)
		}
		cfg := NewConfig()
		err = json.Unmarshal(data, cfg)
		if err != nil {
			logger.Panic("unmarshal config fail", "path", path,
				"data", string(data))
		}
		G = cfg
	})

	return nil
}
