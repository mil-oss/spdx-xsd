package main

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// securityDatastruct ...
	securityDatastruct interface{}
)

func main() {
	securityDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv("config/spdx-security-cfg.json")
	xsdprov.BuildIep(securityDatastruct)
	xsdprov.StartWeb()
}
