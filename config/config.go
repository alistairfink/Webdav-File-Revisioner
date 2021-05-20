package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetConfig() Config {
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return config
}
