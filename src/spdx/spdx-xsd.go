package main

import (
	"xsdprov"
)

func main() {
	xsdprov.Setup("spdx", "", Resources, Resdirectories, NewLicense())
}

// Resources ...
var Resources = map[string]string{
	"ref.xsd":       "/xml/xsd/spdx-ref.xsd",
	"iep.xsd":       "/xml/xsd/spdx-license.xsd",
	"XMLSchema.xsd": "/xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":      "/xml/xsd/ext/w3c/xslt.xsd",
	"iep_xsd.xsl":   "/xml/xsl/spdx_license_iep.xsl",
	//"xml_instance.xsl":         "/xml/xsl/xml_instance.xsl",
	//"xsd_json.xsl":             "/xml/xsl/xsd_json.xsl",
	//"xml_json.xsl":             "/xml/xsl/xml_json.xsl",
	"go-gen.xsl":               "/xml/xsl/go-gen.xsl",
	"go-test-gen.xsl":          "/xml/xsl/go-test-gen.xsl",
	"test_data.xml":            "/xml/instance/test_data.xml",
	"test_instance.xml":        "/xml/instance/test_instance.xml",
	"test_instance-golang.xml": "/xml/instance/test_instance-golang.xml",
	"iep_xsd.json":             "/json/iep_xsd.json",
	"ref_xsd.json":             "/json/ref_xsd.json",
	"test_instance.json":       "/json/test_instance.json",
	"xsd-struct.go":            "/src/spdx/xsd-struct.go",
	"xsd-test.go":              "/src/spdx/xsd-test.go",
	"xsd.go":                   "/src/spdx/xsd.go",
	"provenance_report.json":   "/tests/provenance_report.json",
	"resources.json":           "/json/resources.json",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":     "/xml/xsd/niem",
	"w3c":      "/xml/xsd/w3c",
	"xsl":      "/xml/xsl",
	"instance": "/xml/instance",
	"seva":     "/src/seva",
	"xsdprov":  "/src/xsdprov",
}
