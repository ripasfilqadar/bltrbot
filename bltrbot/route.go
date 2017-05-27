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
	Function    string `yaml:"function"`
	LenArgs     string `yaml:"len_args"`
	Description string `yaml:"description"`
	Admin       bool   `yaml:Admin`
}

var Routes Route

func InitRoute() {
	source, err := ioutil.ReadFile("constant/route.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Routes)
	if err != nil {
		panic(err)
	}
}
