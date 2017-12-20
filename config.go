package minappapi

import "github.com/BurntSushi/toml"

type Config struct {
	ReaderMinApp ReaderMinApp
	Database     Database
}

//ReaderMinApp 配置
type ReaderMinApp struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

//Database 配置
type Database struct {
	Type     string `toml:"type"`
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Sslmode  string `toml:"sslmode"`
	Dbname   string `toml:"dbname"`
	Port     int    `toml:"port"`
}

var confFile = "conf.toml"
var config Config

func init() {
	GetConf()
	DB().AutoMigrate(&Subscribe{}, &Fans{}, &Post{})
}

//GetConf 获取config
func GetConf() Config {
	if config.ReaderMinApp.AppID == "" {
		if _, err := toml.DecodeFile(confFile, &config); err != nil {
			panic(err)
		}
	}
	return config
}
