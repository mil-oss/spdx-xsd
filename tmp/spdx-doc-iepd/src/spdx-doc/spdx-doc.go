package spdxdoc

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
	BuildDocIEP(Config)
}

// BuildDocIEP ...
func BuildDocIEP(config string) {
	SpdxDocDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SpdxDocDatastruct)
	//xsdprov.StartWeb(xsdprov.Homeurl)
}
