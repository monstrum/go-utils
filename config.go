package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func Parse(file string) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, errors.New("missing file")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}