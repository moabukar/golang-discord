package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string `json : "Token"`
	BotPrefix string `json : "BotPrefix`
)

type ConfigStruct struct {
	Token string `json : "Token:`
	BotPrefix string `json : "BotPrefix`
}

func ReadConfig() error {
	fmt.Println("Reading config file")
	file, err := ioioutil.ReadFile("./config.json")

	// error handling
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
}

