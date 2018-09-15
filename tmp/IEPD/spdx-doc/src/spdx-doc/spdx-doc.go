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
	SPDXLicenseDatastruct = NewSpdxDocument()
	xsdprov.InitXSDProv(Resources, Resdirectories, "config/spdx-doc-cfg.json")
	BuildLicenseIep()
	xsdprov.StartWeb()
}

// Resources ...
var Resources = map[string]string{
	"spdx-ref.xsd":                      "xml/xsd/spdx-ref.xsd",
	"spdx-doc.xsd":                      "xml/xsd/spdx-doc.xsd",
	"XMLSchema.xsd":                     "xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":                          "xml/xsd/ext/w3c/xslt.xsd",
	"spdx-doc-iep.xsl":                  "xml/xsl/spdx-doc-iep.xsl",
	"spdx-doc-instance.xsl":             "xml/xsl/spdx-doc-instance.xsl",
	"xsd-json.xsl":                      "xml/xsl/xsd-json.xsl",
	"xml-json.xsl":                      "xml/xsl/xml-json.xsl",
	"go-gen-doc.xsl":                    "xml/xsl/go-gen-doc.xsl",
	"go-gen-doc-test.xsl":               "xml/xsl/go-gen-doc-test.xsl",
	"spdx-test-data.xml":                "xml/instance/spdx-test-data.xml",
	"spdx-doc-test-instance.xml":        "xml/instance/spdx-doc-test-instance.xml",
	"spdx-doc-test-instance-golang.xml": "xml/instance/spdx-doc-test-instance-golang.xml",
	"spdx-ref-xsd.json":                 "json/spdx-ref-xsd.json",
	"spdx-doc-iep-xsd.json":             "json/spdx-doc-xsd.json",
	"spdx-doc-test-instance.json":       "json/spdx-doc-test-instance.json",
	"spdx-doc-struct.go":                "src/spdx-doc/spdx-doc-struct.go",
	"spdx-doc_test.go":                  "src/spdx-doc/spdx-doc_test.go",
	"spdx-xsd.go":                       "src/spdx-doc/spdx-doc.go",
	"provenance-report.json":            "resources/reports/provenance-report.json",
	"resources.json":                    "json/resources.json",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":     "xml/xsd/ext/niem",
	"w3c":      "xml/xsd/ext/w3c",
	"xsl":      "xml/xsl",
	"xsd":      "xml/xsd",
	"json":     "json",
	"instance": "xml/instance",
	"spdx":     "src/spdx-doc",
	"xsdprov":  "src/xsdprov",
}
