package cfg

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type CFG struct {
	Database struct {
		Driver string `yaml:"driver"`
		Path   string `yaml:"path"`
	}
}

var Cfg CFG

var configPath string = "../etc/cfg.yml"

func InitConfig() error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("get error: %v when loading config file: %s \n", err, configPath)
		return err
	}
	err = yaml.Unmarshal(configFile, Cfg)
	if err != nil {
		log.Fatalf("get error: %v when unmarshalling config file: %s \n", err, configPath)
		return err
	}
	return nil
}
