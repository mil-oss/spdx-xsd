package xsdprov

import (
	"encoding/json"
	"io/ioutil"
)

// Cfg ...
type Cfg struct {
	Project     string     `json:"project,omitempty"`
	Configfile  string     `json:"configfile,omitempty"`
	Reflink     string     `json:"reflink,omitempty"`
	Testlink    string     `json:"testlink,omitempty"`
	Homeurl     string     `json:"homeurl,omitempty"`
	Port        string     `json:"port,omitempty"`
	Dbloc       string     `json:"dbloc,omitempty"`
	Tempdir     string     `json:"tempdir,omitempty"`
	Temppath    string     `json:"temppath,omitempty"`
	Resources   []Resource `json:"resources,omitempty"`
	Directories []Resource `json:"directories,omitempty"`
}

//GetConfig ...
func GetConfig(cfgpath string) Cfg {
	jf, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		panic(err)
	}
	var c Cfg
	jerr := json.Unmarshal([]byte(jf), &c)
	check(jerr)
	return c
}
