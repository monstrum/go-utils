package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type MappedString map[string]interface{}

func Parse(file string) interface{} {
	parsed, err := ParseStruct(file)
	if err != nil {
		return nil
	}
	return parsed
}

func ParseMap(file string) (map[string]interface{}, error) {
	parsed, err := ParseStruct(file)
	if err != nil {
		return nil, err
	}

	if brands, ok := parsed.(MappedString); ok {
		return brands, nil
	}
	return nil, errors.New("error")
}

func ParseStruct(file string) (interface{}, error) {
	var result *interface{}
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