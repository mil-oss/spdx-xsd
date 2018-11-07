package spdxsecism

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
	//SecurityIsmIEP(Config)
	//xsdprov.StartWeb(Config, SecurityISMDatastruct)
}

// SecurityIsmIEP ...
func SecurityIsmIEP(config string) {
	SecurityISMDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SecurityISMDatastruct)
}
