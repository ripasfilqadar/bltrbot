package main

import (
	"io/ioutil"
	//	"reflect"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Host    string             `yaml:"host"`
	Command map[string]Command `yaml:"command"`
}

type Command struct {
	Function string `yaml:"function"`
}

var Routes Route

func InitRoute() {
	source, err := ioutil.ReadFile("constant/route.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Routes)
}
