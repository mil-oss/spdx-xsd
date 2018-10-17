package spdxsec

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// SecurityDatastruct ...
	SecurityDatastruct interface{}
)

// BuildSecurityIEP ...
func BuildSecurityIEP(config string) {
	SecurityDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SecurityDatastruct)
	//xsdprov.StartWeb(xsdprov.Homeurl)
}
