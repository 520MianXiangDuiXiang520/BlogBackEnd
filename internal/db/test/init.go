package test

import (
	"JuneBlog/internal/config"
	"JuneBlog/internal/db"
)

func init() {
	err := config.InitConfig("./config.json")
	if err != nil {
		panic("init cfg error: " + err.Error())
	}
	err = db.InitDatabase()
	if err != nil {
		panic("init db error: " + err.Error())
	}
}
