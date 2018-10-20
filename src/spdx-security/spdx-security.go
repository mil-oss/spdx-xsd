package spdxsec

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
	//BuildSecurityIEP(Config)
	//xsdprov.StartWeb(Config, SecurityDatastruct)
}

// BuildSecurityIEP ...
func BuildSecurityIEP(config string) {
	SecurityDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SecurityDatastruct)
}
