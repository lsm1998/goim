package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	configName = "application.yml"
)

func ScanConfig(c interface{}) error {
	yamlFile, err := ioutil.ReadFile(configName)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}
	return nil
}
