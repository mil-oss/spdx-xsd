package main

import (
	"xsdprov"
)

var testinstances = map[string]string{
	"spdx-license-test-instance.xml": "spdx-license-test-instance.xml",
	"spdx-doc-test-instance.xml":     "xml/instance/spdx-doc-test-instance.xml",
}

func main() {
	xsdprov.LicenceDatastruct = NewLicense()
	xsdprov.SPDXDocDatastruct = NewSpdxDocument()
	xsdprov.Setup(Resources, Resdirectories)
}

// Resources ...
var Resources = map[string]string{
	"spdx-ref.xsd":                          "xml/xsd/spdx-ref.xsd",
	"spdx-license.xsd":                      "xml/xsd/spdx-license.xsd",
	"spdx-doc.xsd":                          "xml/xsd/spdx-doc.xsd",
	"XMLSchema.xsd":                         "xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":                              "xml/xsd/ext/w3c/xslt.xsd",
	"spdx-license-iep.xsl":                  "xml/xsl/spdx-license-iep.xsl",
	"spdx-doc-iep.xsl":                      "xml/xsl/spdx-doc-iep.xsl",
	"spdx-license-instance.xsl":             "xml/xsl/spdx-license-instance.xsl",
	"spdx-doc-instance.xsl":                 "xml/xsl/spdx-doc-instance.xsl",
	"xsd-json.xsl":                          "xml/xsl/xsd-json.xsl",
	"xml-json.xsl":                          "xml/xsl/xml-json.xsl",
	"go-gen-lic.xsl":                        "xml/xsl/go-gen-lic.xsl",
	"go-gen-lic-test.xsl":                   "xml/xsl/go-gen-lic-test.xsl",
	"go-gen-doc.xsl":                        "xml/xsl/go-gen-doc.xsl",
	"go-gen-doc-test.xsl":                   "xml/xsl/go-gen-doc-test.xsl",
	"spdx-test-data.xml":                    "xml/instance/spdx-test-data.xml",
	"spdx-license-test-instance.xml":        "xml/instance/spdx-license-test-instance.xml",
	"spdx-doc-test-instance.xml":            "xml/instance/spdx-doc-test-instance.xml",
	"spdx-license-test-instance-golang.xml": "xml/instance/spdx-license-test-instance-golang.xml",
	"spdx-doc-test-instance-golang.xml":     "xml/instance/spdx-doc-test-instance-golang.xml",
	"spdx-ref-xsd.json":                     "json/spdx-ref-xsd.json",
	"spdx-license-iep-xsd.json":             "json/spdx-license-xsd.json",
	"spdx-doc-iep-xsd.json":                 "json/spdx-doc-xsd.json",
	"spdx-license-test-instance.json":       "json/spdx-license-test-instance.json",
	"spdx-doc-test-instance.json":           "json/spdx-doc-test-instance.json",
	"spdx-license-struct.go":                "src/spdx/spdx-license-struct.go",
	"spdx-license_test.go":                  "src/spdx/spdx-license_test.go",
	"spdx-doc-struct.go":                    "src/spdx/spdx-doc-struct.go",
	"spdx-doc_test.go":                      "src/spdx/spdx-doc_test.go",
	"spdx-xsd.go":                           "src/spdx/spdx-xsd.go",
	"provenance-report.json":                "resources/reports/provenance-report.json",
	"resources.json":                        "json/resources.json",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":     "xml/xsd/ext/niem",
	"w3c":      "xml/xsd/ext/w3c",
	"xsl":      "xml/xsl",
	"xsd":      "xml/xsd",
	"json":     "json",
	"instance": "xml/instance",
	"licences": "resources/xml-licenses",
	"spdx":     "src/spdx",
	"xsdprov":  "src/xsdprov",
}
