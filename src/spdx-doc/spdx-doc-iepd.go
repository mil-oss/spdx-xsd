package main

import (
	"fmt"
	"log"
	"time"
	"xsdprov"
)

var (
	dbloc        string
	temppath     string
	name         string
	reflink      string
	testlink     string
	port         string
	resources    map[string]string
	resourcedirs map[string]string
	tempfiles    = map[string]string{}
	tempdigests  = map[string]string{}
)

//SPDXLicenseDatastruct ...
var SPDXLicenseDatastruct interface{}

//BuildLicenseIep ... Generate XML, Code and Test Artifacts
func BuildLicenseIep() (map[int64]xsdprov.ProvEntry, []error, error) {
	log.Println("BuildIep")
	log.Println("reflink " + reflink)
	xsdprov.GetSourceResources("spdx-ref.xsd", "spdx-test-data.xml")
	generateResources()
	validateResources()
	xsdprov.ProvenanceRpt()
	return xsdprov.Provreport, valerr, iepderr
}

func generateResources() (err error) {
	log.Println("Generate Resources")
	//spdx-doc.xsd - Information Exchange Package XML Schema
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("spdx-doc-iep.xsl", "spdx-ref.xsd", "spdx-doc.xsd"))
	check(err)
	//spdx-doc-test-instance.xml - Information Exchange Package XML Instance
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("spdx-doc-instance.xsl", "spdx-doc.xsd", "spdx-doc-test-instance.xml"))
	check(err)
	//JSON
	//spdx-ref-xsd.json - JSON representation of spdx-ref.xsd
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("xsd-json.xsl", "spdx-ref.xsd", "spdx-ref-xsd.json"))
	check(err)
	//spdx-doc-iep-xsd.json - JSON representation of spdx-doc-iep-xsd
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("xsd-json.xsl", "spdx-doc.xsd", "spdx-doc-iep-xsd.json"))
	check(err)
	//spdx-doc-test-instance.json - JSON representation doc-test-instance.xml
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("xml-json.xsl", "spdx-doc-test-instance.xml", "spdx-doc-test-instance.json"))
	check(err)
	//spdx-doc-struct.go - Golang struct iep.go
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("go-gen-doc.xsl", "spdx-doc.xsd", "spdx-doc-struct.go"))
	check(err)
	//spdx-doc_test.go - Golang test iep.go
	xsdprov.Provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource(getemppaths("go-gen-doc-test.xsl", "spdx-doc.xsd", "spdx-doc_test.go"))
	check(err)
	//Marshal instance
	xsdprov.Provreport[time.Now().UnixNano()] = xsdprov.MarshalXML(temppath+Resources["spdx-doc-test-instance.xml"], Resources["spdx-doc-test-instance-golang.xml"], SPDXLicenseDatastruct)
	return err
}

func getemppaths(xslname string, xmlname string, resultname string) (string, string, string) {
	var xslpath = temppath + Resources[xslname]
	var xmlpath = temppath + Resources[xmlname]
	var resultemppath = temppath + Resources[resultname]
	if val, ok := tempfiles[xslname]; ok {
		xslpath = val
	}
	if val, ok := tempfiles[xmlname]; ok {
		xmlpath = val
	}
	log.Println("xslpath: " + xslpath)
	log.Println("xmlpath: " + xmlpath)
	log.Println("resultemppath: " + resultemppath)
	return xslpath, xmlpath, resultemppath
}

func validateResources() (err error) {
	log.Println("Validate Resources")
	var errs []error
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-ref.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-doc.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-doc-test-instance.xml", "spdx-doc.xsd")
	check(err)
	checka(errs)
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-doc-test-instance.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-doc-test-instance-golang.xml", "spdx-doc.xsd")
	check(err)
	checka(errs)
	xsdprov.Provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-doc-test-instance-golang.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	return err
}

func check(e error) error {
	if e != nil {
		fmt.Printf("error: %v\n", e)
	}
	iepderr = e
	return e
}

func checka(e []error) []error {
	if e != nil {
		fmt.Printf("error: %v\n", e)
	}
	valerr = e
	return e
}
