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
	Config = "config/spdx-doc-cfg.json"
	// SpdxDocDatastruct ...
	SpdxDocDatastruct interface{}
)

func main() {
	xsdprov.InitXSDProv(Config)
	//BuildDocIEP()
	xsdprov.StartWeb(Config, SpdxDocDatastruct)
}

// BuildDocIEP ...
func BuildDocIEP() {
	SpdxDocDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv(Config)
	xsdprov.BuildIep(SpdxDocDatastruct)
	//xsdprov.StartWeb(xsdprov.Homeurl)
}
