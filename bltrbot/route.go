package main

import (
	"io/ioutil"
	//	"reflect"

	//	"encoding/json"

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
	Scope       string `yaml:scope`
}

func (c *Command) IsUser() bool {
	return c.Scope == "user"
}

func (c *Command) IsGroup() bool {
	return c.Scope == "group"
}

func (c *Command) IsAdmin() bool {
	return c.Scope == "admin"
}

type CallbackMessage struct {
	Controller string `json:"controller"`
	Data       string `json:"data"`
	MessageId  string `json:"message_id"`
}

var Routes Route

var CallbackMsg CallbackMessage

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
