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

func main() {
	SpdxDocDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv("config/spdx-doc-cfg.json")
	xsdprov.BuildIep(SpdxDocDatastruct)
	xsdprov.StartWeb()
}
