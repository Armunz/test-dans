package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

func ReadConfig() (result Config, err error) {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	jsonFile, err := os.Open(*fptr)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &result)

	return
}
