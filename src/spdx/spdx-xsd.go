package main

import (
	"xsdprov"
)

func main() {
	xsdprov.Setup("spdx", "", Resources, Resdirectories, NewLicense())
}

// Resources ...
var Resources = map[string]string{
	"ref.xsd":                  "resources/xml/xsd/spdx-ref.xsd",
	"iep.xsd":                  "resources/xml/xsd/spdx-license.xsd",
	"XMLSchema.xsd":            "resources/xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":                 "resources/xml/xsd/ext/w3c/xslt.xsd",
	"iep_xsd.xsl":              "resources/xml/xsl/spdx_license_iep.xsl",
	"xml_instance.xsl":         "resources/xml/xsl/xml_instance.xsl",
	"xsd_json.xsl":             "resources/xml/xsl/xsd_json.xsl",
	"xml_json.xsl":             "resources/xml/xsl/xml_json.xsl",
	"go-gen.xsl":               "resources/xml/xsl/go-gen.xsl",
	"go-test-gen.xsl":          "resources/xml/xsl/go-test-gen.xsl",
	"test_data.xml":            "resources/xml/instance/test_data.xml",
	"test_instance.xml":        "resources/xml/instance/test_instance.xml",
	"test_instance-golang.xml": "resources/xml/instance/test_instance-golang.xml",
	"iep_xsd.json":             "resources/json/spdx-license.json",
	"ref_xsd.json":             "resources/json/spdx-ref.json",
	"test_instance.json":       "resources/json/test_instance.json",
	"xsd-struct.go":            "src/spdx/spdx-license.go",
	"xsd-test.go":              "src/spdx/spdx-license_test.go",
	"spdx-xsd.go":              "src/spdx/spdx-xsd.go",
	"provenance_report.json":   "resources/tests/provenance_report.json",
	"resources.json":           "resources/json/resources.json",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":     "resources/xml/xsd/ext/niem",
	"w3c":      "resources/xml/xsd/ext/w3c",
	"xsl":      "resources/xml/xsl",
	"xsd":      "resources/xml/xsd",
	"json":     "resources/json",
	"instance": "resources/xml/instance",
	"licences": "resources/xml-licenses",
	"spdx":     "src/spdx",
	"xsdprov":  "src/xsdprov",
}
