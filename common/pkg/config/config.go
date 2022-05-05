package config

import (
	"os"
	"sync"
	"gitee.com/zince/micro-service/common/pkg/types"
	"time"
)

var once *sync.Once

type Db struct {
	Address               string `json:"address"`
	Database              string `json:"database"`
	User                  string `json:"user"`
	Password              string `json:"password"`
	Charset               string `json:"charset"`
	MaxConnections        int    `json:"max_connections"`
	MaxIdeConnections     int    `json:"max_ide_connections"`
	ConnectionMaxLifeTime time.Duration    `json:"connection_max_life_time"`
}

type Configuration struct {
	Db *Db `json:"db"`
}

var config *Configuration

func LoadConfig() *Configuration  {
	once.Do(func() {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		database := os.Getenv("DB_DATABASE")
		password := os.Getenv("DB_PASSWORD")
		dbMaxConnections, _ := types.StringToInt(os.Getenv("DB_MAX_CONNECTIONS"))
		dbMaxIdeConnections, _ := types.StringToInt(os.Getenv("DB_MAX_IDE_CONNECTIONS"))
		dbConnectionMaxLifeTime, _ := types.StringToInt(os.Getenv("DB_CONNECTIONS_MAX_LIFE_TIME"))

		dbconfig := &Db{
			Address:               host,
			Database:              database,
			User:                  user,
			Password:              password,
			Charset:               "utf8",
			MaxConnections:        dbMaxConnections,
			MaxIdeConnections:     dbMaxIdeConnections,
			ConnectionMaxLifeTime: time.Duration(dbConnectionMaxLifeTime) * time.Minute,
		}
		config = &Configuration{Db:dbconfig}
	})
	return config
}