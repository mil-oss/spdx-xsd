package xsdprov

import (
	"encoding/json"
	"io/ioutil"
)
const (
	cfgpath string = "config/xsdccm.json"
)


//Cfg ...
type Cfg struct {
	Project string `json:"project,omitempty"`
	Reflink      = `json:"reflink,omitempty"`
	Testlink     = `json:"testlink,omitempty"`
}

func getConfig() Cfg {
	jf, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		panic(err)
	}
	var c Cfg
	jerr := json.Unmarshal([]byte(jf), &c)
	check(jerr)
	return c
}
