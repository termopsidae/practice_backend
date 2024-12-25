package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Conf struct {
	BSCManagerInfo struct {
		Dial            string `yaml:"dial"`
		ContractAddress string `yaml:"contractAddress"`
		Pk              string `yaml:"pk"`
		PkAD            string `yaml:"pkAD"`
	} `yaml:"bscManagerInfo"`
}

func (m *Conf) GetConf() *Conf {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error")
	}
	index := strings.Index(basePath, "degen-backen")
	fileName := filepath.Join(basePath[0:index+12], "conf.yaml")
	//fmt.Println("config:", fileName)
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("load conf error")
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	return m
}
