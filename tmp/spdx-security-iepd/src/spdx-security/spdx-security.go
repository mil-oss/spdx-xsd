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
	Config = "config/spdx-security-cfg.json"
	// SecurityDatastruct ...
	SecurityDatastruct interface{}
)

func main() {
	xsdprov.InitXSDProv(Config)
	//BuildSecurityIEP()
	xsdprov.StartWeb(Config, SecurityDatastruct)
}

// BuildSecurityIEP ...
func BuildSecurityIEP() {
	SecurityDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(Config)
	xsdprov.BuildIep(SecurityDatastruct)
}
