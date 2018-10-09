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
	xsdprov.InitXSDProv(Resources, Sources, Resdirectories, "config/spdx-seva-cfg.json")
	xsdprov.BuildIep(SevaDatastruct)
	xsdprov.StartWeb()
}

// Sources ...
var Sources = map[string]string{
	"refxsd":               "xml/xsd/spdx-seva-ref.xsd",
	"iepxsd":               "IEPD/spdx-seva/xml/xsd/spdx-seva-iep.xsd",
	"xmlschemaxsd":         "xml/xsd/ext/w3c/XMLSchema.xsd",
	"xsltxsd":              "xml/xsd/ext/w3c/xslt.xsd",
	"iepxsdxsl":            "IEPD/spdx-seva/xml/xsl/spdx-seva-iep.xsl",
	"instancexsl":          "IEPD/spdx-seva/xml/xsl/spdx-seva_instance.xsl",
	"gogenxsdxsl":          "IEPD/spdx-seva/xml/xsl/spdx-seva-go-gen.xsl",
	"gotestgenxsl":         "IEPD/spdx-seva/xml/xsl/spdx-seva-go-test-gen.xsl",
	"testdataxml":          "IEPD/spdx-seva/xml/instance/spdx-seva-test_data.xml",
	"instancexml":          "IEPD/spdx-seva/xml/instance/spdx-seva-instance.xml",
	"instancegolangxml":    "IEPD/spdx-seva/xml/instance/test_instance-golang.xml",
	"structgo":             "src/spdx-seva/spdx-seva-struct.go",
	"structtestgo":         "src/spdx-seva/spdx-seva-struct_test.go",
	"maingo":               "src/spdx-seva/spdx-seva.go",
	"iepxsdjson":           "IEPD/spdx-seva/json/iep_xsd.json",
	"refxsdjson":           "IEPD/spdx-seva/json/ref_xsd.json",
	"instancejson":         "IEPD/spdx-seva/json/test-instance.json",
	"provenancereportjson": "IEPD/spdx-seva/json/provenance_report.json",
	"resourcesjson":        "IEPD/spdx-seva/json/resources.json",
	"iepxsl":               "xml/xsl/iep.xsl",
	"xmlinstancexsl":       "xml/xsl/xml-instance.xsl",
	"gogenxsl":             "xml/xsl/go-gen.xsl",
	"xsdjsonxsl":           "xml/xsl/xsd-json.xsl",
	"xmljsonxsl":           "xml/xsl/xml-json.xsl",
}

// Resources ...
var Resources = map[string]string{
	"refxsd":               "xml/xsd/spdx-seva-ref.xsd",
	"iepxsd":               "xml/xsd/spdx-seva-iep.xsd",
	"xmlschemaxsd":         "xml/xsd/ext/w3c/XMLSchema.xsd",
	"xsltxsd":              "xml/xsd/ext/w3c/xslt.xsd",
	"iepxsdxsl":            "xml/xsl/spdx-seva-iep.xsl",
	"instancexsl":          "xml/xsl/spdx-seva_instance.xsl",
	"gogenxsdxsl":          "xml/xsl/spdx-seva-go-gen.xsl",
	"gotestgenxsl":         "xml/xsl/spdx-seva-go-test-gen.xsl",
	"testdataxml":          "xml/instance/spdx-seva-test_data.xml",
	"instancexml":          "xml/instance/spdx-seva-instance.xml",
	"instancegolangxml":    "xml/instance/spdx-seva-instance-golang.xml",
	"structgo":             "src/spdx-seva/spdx-seva-struct.go",
	"structtestgo":         "src/spdx-seva/spdx-seva-struct_test.go",
	"maingo":               "src/spdx-seva/spdx-seva.go",
	"iepxsdjson":           "json/spdx-seva-iep-xsd.json",
	"refxsdjson":           "json/spdx-seva-ref-xsd.json",
	"instancejson":         "json/spdx-seva-test-instance.json",
	"provenancereportjson": "json/spdx-seva-provenance_report.json",
	"resourcesjson":        "json/spdx-seva-resources.json",
	"iepxsl":               "xml/xsl/common/iep.xsl",
	"xmlinstancexsl":       "xml/xsl/common/xml-instance.xsl",
	"gogenxsl":             "xml/xsl/common/go-gen.xsl",
	"xsdjsonxsl":           "xml/xsl/common/xsd-json.xsl",
	"xmljsonxsl":           "xml/xsl/common/xml-json.xsl",
	"zipiepd":              "spdx-seva-iepd.zip",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":      "xml/xsd/ext/niem",
	"w3c":       "xml/xsd/ext/w3c",
	"spdx-seva": "src/spdx-seva",
	"xsdprov":   "src/xsdprov",
}
