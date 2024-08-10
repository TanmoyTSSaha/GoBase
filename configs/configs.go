package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Config Configuration

type Configuration struct {
	Server ServerConfig `yaml:"server"`
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Internals Internals `yaml:"internals"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
}

type MongoDBConfig struct {
	URI string `yaml:"uri"`
	Database string `yaml:"database"`
}

type Internals struct {
	IsLogDependenciesCreated bool `yaml:"isLogDependenciesCreated"`
}

func LoadConfig() error {
	file, err := os.Open("configs/configs.yaml")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	decodeErr := decoder.Decode(&Config)
	if decodeErr != nil {
		return decodeErr
	}

	return nil
}

func UpdateCongif(config *Configuration) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	
	err = os.WriteFile("configs/configs.yaml", data, 0644)
	if err != nil {
		return err
	}

	return nil
}