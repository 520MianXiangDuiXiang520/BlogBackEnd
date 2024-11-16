package env

import (
	"os"
	"path"
	"sync"
)

type GlobalEnv struct {
	CfgPath string
	LogPath string
}

func getOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultValue
}

var (
	once              sync.Once
	G                 *GlobalEnv
	defaultDataPath   = "/data/juneblog"
	defaultConfigPath = path.Join(defaultDataPath, "config/config.json")
	defaultLogPath    = path.Join(defaultDataPath, "log")
)

func InitGlobalEnv() {
	once.Do(func() {
		G = &GlobalEnv{
			CfgPath: getOrDefault("CFG_PATH", defaultConfigPath),
			LogPath: getOrDefault("LOG_PATH", defaultLogPath),
		}
	})
}
