package spdxdoc

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
	// SpdxDocDatastruct ...
	SpdxDocDatastruct interface{}
)

// BuildDocIEP ...
func BuildDocIEP(config string) {
	SpdxDocDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SpdxDocDatastruct)
	xsdprov.StartWeb(xsdprov.Homeurl)
}
