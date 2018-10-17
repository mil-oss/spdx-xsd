package main

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// Config ...
	Config = "config/spdx-sec-ism-cfg.json"
	// SecurityISMDatastruct ...
	SecurityISMDatastruct interface{}
)

func main() {
	xsdprov.InitXSDProv(Config)
	//SecurityIsmIEP()
	xsdprov.StartWeb(Config, SecurityISMDatastruct)
}

// SecurityIsmIEP ...
func SecurityIsmIEP() {
	SecurityISMDatastruct = NewSoftwareEvidenceArchiveISM()
	xsdprov.InitXSDProv(Config)
	xsdprov.BuildIep(SecurityISMDatastruct)
}
