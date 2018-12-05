package config

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobestsdk/gobase/log"
	"io/ioutil"
)

type Config struct {
	HttpPort     int    `json:"http_port"`
	FsPort       int    `json:"fs_port"`
	Object4dPort int    `json:"object4d_port"`
	Mysql        string `json:"mysql"`
	Redis        struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
}

var (
	APPConfig Config = Config{
		HttpPort: 30000,
		Mysql:    "",
	}
)

func ParseConfig(configfilepath string) error {
	data, err := ioutil.ReadFile(configfilepath)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file read "})
		return err
	}

	err = json.Unmarshal([]byte(data), &APPConfig)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file parse "})
		return err
	}
	log.Info(log.Fields{"app": "config file", "config": APPConfig})
	return nil
}
