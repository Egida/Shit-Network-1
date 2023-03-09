package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"
)

type Config struct {
	Server    string
	Port      string
	BotServer string
	BotPort   string

	AntiDuplicate bool
	BotToken      string
	ChatId        string
	Logging       bool
}

func GetConfig() *Config {

	var config Config

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		color.HiRed("Can't read config")
		return nil
	}

	json.Unmarshal(file, &config)

	return &config

}

func GetBlacklist() string {

	file, err := ioutil.ReadFile("./data/blacklist.txt")

	if err != nil {
		color.HiRed("Can't read blacklist file")
		return ""
	}

	return string(file)

}
