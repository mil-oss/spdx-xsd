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
	xsdprov.InitXSDProv(Config)
	//BuildDocIEP()
	//xsdprov.StartWeb(Config, SpdxDocDatastruct)
}

// BuildDocIEP ...
func BuildDocIEP(config string) {
	SpdxDocDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv(config)
	xsdprov.BuildIep(SpdxDocDatastruct)
}
