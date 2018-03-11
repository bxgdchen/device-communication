package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// C config info
var C Config

// Config ...
type Config struct {
	Stations []Station `yaml:"stations,flow"`
}

// Station ...
type Station struct {
	Name     string   `yaml:"name"`
	APIKey   string   `yaml:"api_key"`
	Token    string   `yaml:"token"`
	ClientID string   `yaml:"client_id"`
	IsGather bool     `yaml:"is_gather"`
	Devices  []Device `yaml:"devices,flow"`
}

// Device ...
type Device struct {
	Name           string  `yaml:"name"`
	IsCalculate    bool    `yaml:"is_calculate"`
	Ins            string  `yaml:"ins"`
	InsMode        string  `yaml:"ins_mode"`
	InsFrequency   int     `yaml:"ins_frequency"`
	Order          string  `yaml:"order"`
	Format         string  `yaml:"format"`
	ConversionRate float64 `yaml:"conversion_rate"`
	Plus           float64 `yaml:"plus"`
	Factor         float64 `yaml:"factor"`
	Prefix         string  `yaml:"prefix"`
}

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("read config file fail, %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &C)
	if err != nil {
		log.Fatalf("parse config file fail, %v", err)
	}
}
