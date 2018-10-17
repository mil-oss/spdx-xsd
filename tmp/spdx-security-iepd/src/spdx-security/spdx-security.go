package spdxsec

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// SecurityURL ...
	SecurityURL string
	// SecurityDatastruct ...
	SecurityDatastruct interface{}
)

// BuildSecurityIEP ...
func BuildSecurityIEP(config string) {
	SecurityDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SecurityDatastruct)
	cfg := xsdprov.GetConfig(config)
	SecurityURL = cfg.Homeurl
	//xsdprov.StartWeb(xsdprov.Homeurl)
}
