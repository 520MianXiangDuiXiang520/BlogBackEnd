package config

import (
	"JuneBlog/patch/logger"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	*Common
	Db       DbConfig `json:"db"`
	CORSList []string `json:"cors_list"`
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
			fmt.Println(err)
			logger.Panic("read config fail", "path", path, "err", err)
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
