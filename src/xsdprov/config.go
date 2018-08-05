package xsdprov

import (
	"encoding/json"
	"io/ioutil"
)

var (
	provreport = map[int64]ProvEntry{}
)

// Cfg ...
type Cfg struct {
	Project  string `json:"project,omitempty"`
	Reflink  string `json:"reflink,omitempty"`
	Testlink string `json:"testlink,omitempty"`
	Homeurl  string `json:"homeurl,omitempty"`
	Port     string `json:"port,omitempty"`
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
