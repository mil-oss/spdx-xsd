package main

import (
	"encoding/json"
	"io/ioutil"
)

//Cfg ...
type Cfg struct {
	Pckg string `json:"pckg,omitempty"`
	Root string `json:"root,omitempty"`
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

func getConfig() Cfg {
	jf, err := ioutil.ReadFile("config/xsdccm.json")
	if err != nil {
		panic(err)
	}
	var c Cfg
	jerr := json.Unmarshal([]byte(jf), &c)
	check(jerr)
	return c
}
