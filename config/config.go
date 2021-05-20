package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetConfig() Config {
	file, err := os.Open("./config.json")
	if err != nil {
		// TODO: USE EMAILER
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		// TODO: USE EMAILER
		panic(err)
	}

	var config Config
	json.Unmarshal(bytes, &config)
	println(config.FilePath, config.DestinationPath)
	return config
}
