package main

import (
	"xsdprov"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]xsdprov.ProvEntry{}
)

func main() {
	//var SevaDatastruct = NewSoftwareEvidenceArchive()
	xsdprov.InitXSDProv(Resources, Resdirectories, "config/spdx-doc-cfg.json")
	xsdprov.StartWeb()
}

// Resources ...
var Resources = map[string]string{
	"spdx-seva-ref.xsd":         "IEPD/spdx-seva/xml/xsd/spdx-seva-ref.xsd",
	"spdx-seva-iep.xsd":         "IEPD/spdx-seva/xml/xsd/spdx-seva-iep.xsd",
	"XMLSchema.xsd":             "IEPD/spdx-seva/xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":                  "IEPD/spdx-seva/xml/xsd/ext/w3c/xslt.xsd",
	"spdx-seva-iep.xsl":         "IEPD/spdx-seva/xml/xsl/spdx-seva-iep.xsl",
	"spdx-seva_instance.xsl":    "IEPD/spdx-seva/xml/xsl/spdx-seva_instance.xsl",
	"spdx-seva-go-gen.xsl":      "IEPD/spdx-seva/xml/xsl/spdx-seva-go-gen.xsl",
	"spdx-seva-go-test-gen.xsl": "IEPD/spdx-seva/xml/xsl/spdx-seva-go-test-gen.xsl",
	"spdx-seva-test_data.xml":   "IEPD/spdx-seva/xml/instance/spdx-seva-test_data.xml",
	"spdx-seva-instance.xml":    "IEPD/spdx-seva/xml/instance/spdx-seva-instance.xml",
	"test_instance-golang.xml":  "IEPD/spdx-seva/xml/instance/test_instance-golang.xml",
	"xsd-struct.go":             "IEPD/spdx-seva/src/xsd-struct.go",
	"xsd-test.go":               "IEPD/spdx-seva/src/xsd-test.go",
	"xsd.go":                    "IEPD/spdx-seva/src/xsd.go",
	"iep_xsd.json":              "IEPD/spdx-seva/json/iep_xsd.json",
	"ref_xsd.json":              "IEPD/spdx-seva/json/ref_xsd.json",
	"test_instance.json":        "IEPD/spdx-seva/json/test_instance.json",
	"provenance_report.json":    "IEPD/spdx-seva/json/provenance_report.json",
	"resources.json":            "IEPD/spdx-seva/json/resources.json",
	"iep.xsl":                   "IEPD/spdx-seva/xml/xsl/common/iep.xsl",
	"xml-instance.xsl":          "IEPD/spdx-seva/xml/xsl/common/xml-instance.xsl",
	"go-gen.xsl":                "IEPD/spdx-seva/xml/xsl/common/go-gen.xsl",
	"xsd-json.xsl":              "IEPD/spdx-seva/xml/xsl/common/xsd-json.xsl",
	"xml-json.xsl":              "IEPD/spdx-seva/xml/xsl/common/xml-json.xsl",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":      "xml/xsd/ext/niem",
	"w3c":       "xml/xsd/ext/w3c",
	"xsl":       "IEPD/spdx-seva/xml/xsl",
	"instance":  "IEPD/spdx-seva/xml/instance",
	"spdx-seva": "IEPD/spdx-seva/src/spdx-seva",
	"xsdprov":   "IEPD/spdx-seva/src/xsdprov",
}
