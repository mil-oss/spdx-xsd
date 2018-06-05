package main

import "xsdprov"

//Tpath ...
var Tpath = "/tmp/IEPD/iepd/"

func main() {
	xsdprov.Setup(Tpath, Resources, Resdirectories, NewLicense())
	MakeLicenses(Resources, Tpath)
	xsdprov.BuildIep()
	xsdprov.StartWeb(Tpath)
}

// Resources ...
var Resources = map[string]string{
	"ref.xsd":                  "xml/xsd/spdx-ref.xsd",
	"iep.xsd":                  "xml/xsd/spdx-license.xsd",
	"XMLSchema.xsd":            "xml/xsd/ext/w3c/XMLSchema.xsd",
	"xslt.xsd":                 "xml/xsd/ext/w3c/xslt.xsd",
	"iep_xsd.xsl":              "xml/xsl/spdx_license_iep.xsl",
	"xml_instance.xsl":         "xml/xsl/xml_instance.xsl",
	"xsd_json.xsl":             "xml/xsl/xsd_json.xsl",
	"xml_json.xsl":             "xml/xsl/xml_json.xsl",
	"go-gen.xsl":               "xml/xsl/go-gen.xsl",
	"go-test-gen.xsl":          "xml/xsl/go-test-gen.xsl",
	"test_data.xml":            "xml/instance/test_data.xml",
	"test_instance.xml":        "xml/instance/test_instance.xml",
	"test_instance-golang.xml": "xml/instance/test_instance-golang.xml",
	"iep_xsd.json":             "json/spdx-license.json",
	"ref_xsd.json":             "json/spdx-ref.json",
	"test_instance.json":       "json/test_instance.json",
	"xsd-struct.go":            "src/spdx/spdx-license.go",
	"xsd-test.go":              "src/spdx/spdx-license_test.go",
	"spdx-xsd.go":              "src/spdx/spdx-xsd.go",
	"provenance_report.json":   "resources/reports/provenance_report.json",
	"resources.json":           "json/resources.json",
	"make_license.xsl":         "xml/xsl/make_license.xsl",
	"licenses.rdf":             "resources/licenses.rdf",
	"licenses.json":            "resources/licenses.json",
}

// Resdirectories ...
var Resdirectories = map[string]string{
	"niem":      "xml/xsd/ext/niem",
	"w3c":       "xml/xsd/ext/w3c",
	"xsl":       "xml/xsl",
	"xsd":       "xml/xsd",
	"json":      "json",
	"instance":  "xml/instance",
	"resources": "resources",
	"spdx":      "src/spdx",
	"xsdprov":   "src/xsdprov",
}
