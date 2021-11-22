package configs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Content struct {
	Source Source `yaml:"source"`
	Sink   Sink   `yaml:"sink"`
}

type Source struct {
	Gin struct {
		Enable bool `yaml:"enable"`
		Port   int  `yaml:"port"`
	}
	Rdbms struct {
		Schema string `yaml:"schema"`
		Args   string `yaml:"args"`
	}
}

type Sink struct {
	Mongodb struct {
		Enable         bool   `yaml:"enable"`
		Addrs          string `yaml:"addrs"`
		DbName         string `yaml:"db-name"`
		CollectionName string `yaml:"collection-name"`
	}
	Kafka struct {
		Enable      bool   `yaml:"enable"`
		BrokerAddrs string `yaml:"broker-addrs"`
	}
}

func GetConfig(configPath string) Content {
	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(fmt.Sprintf("yamlFile.Get err #%v ", err))
	}

	var content Content

	err = yaml.Unmarshal(yamlFile, &content)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal: %v", err))
	}

	return content
}
