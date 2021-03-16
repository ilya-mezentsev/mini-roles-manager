package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func mustParse(settingsFilePath string) settings {
	data, err := ioutil.ReadFile(settingsFilePath)
	if err != nil {
		panic(fmt.Errorf("unable to read settings file: %v", err))
	}

	var s settings
	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal settings to struct: %v", err))
	}

	return s
}
