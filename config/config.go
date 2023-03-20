package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadConfig() (result Config, err error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &result)

	return
}
