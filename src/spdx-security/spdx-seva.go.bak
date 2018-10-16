package main

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// SevaDatastruct ...
	SevaDatastruct interface{}
)

func main() {
	SevaDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv("config/spdx-seva-cfg.json")
	xsdprov.BuildIep(SevaDatastruct)
	xsdprov.StartWeb()
}
