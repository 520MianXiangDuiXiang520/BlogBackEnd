package config

type MgoConfig struct {
	Uri           string   `json:"uri"`
	Compressors   []string `json:"compressors"`
	ConnTimeout   int      `json:"conn_timeout"`
	MaxConnecting uint64   `json:"max_connecting"`
	MaxPoolSize   uint64   `json:"max_pool_size"`
	MinPoolSize   uint64   `json:"min_pool_size"`
}

type DbConfig struct {
	Mgo MgoConfig `json:"mongo"`
}
