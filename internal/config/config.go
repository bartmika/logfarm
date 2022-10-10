package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server  serverConf
	DB      dbConf
	Setting settingConf
}

type serverConf struct {
	Port      string `env:"LOGFARM_PORT,required"`
	IP        string `env:"LOGFARM_IP,required"`
	SecretKey []byte `env:"LOGFARM_SECRET_KEY,required"`
}

type dbConf struct {
	FilePath          string `env:"LOGFARM_DB_FILEPATH,required"`
	HasAutoMigrations bool   `env:"LOGFARM_DB_HAS_AUTO_MIGRATIONS,default=true"`
}

type settingConf struct {
	MaxDayAge int `env:"LOGFARM_SETTING_MAX_DAY_AGE,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
