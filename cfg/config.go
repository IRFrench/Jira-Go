package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(filePath string) (Config, error) {
	var newConfig Config

	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return newConfig, err
	}

	if err := yaml.Unmarshal(configFile, &newConfig); err != nil {
		return newConfig, err
	}

	return newConfig, nil
}
