package main

import (
	"fmt"
	"log"
	"os"
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
	tempfiles    map[string]string
	tempdigests  map[string]string
)

//SPDXLicenseDatastruct ...
var SPDXLicenseDatastruct interface{}

//SetupSPDXLicense ...
func SetupSPDXLicense(resrces map[string]string, dirs map[string]string) {
	resources = resrces
	resourcedirs = dirs
	cfg := xsdprov.GetConfig("config/spdx-license-cfg.json")
	dbloc = "/tmp/" + cfg.Project
	temppath = "/tmp/IEPD/spdx-license/"
	name = cfg.Project
	reflink = cfg.Reflink
	testlink = cfg.Testlink
	port = cfg.Port
	err := os.MkdirAll(dbloc+"/db", 0777)
	if err != nil {
		return
	}
	db, err := xsdprov.DbSetup(dbloc)
	check(err)
	e := xsdprov.TempDir(db)
	check(e)
	xsdprov.DirSetup()
	BuildLicenseIep()
	xsdprov.StartWeb(temppath)
}

//BuildLicenseIep ... Generate XML, Code and Test Artifacts
func BuildLicenseIep() (map[int64]xsdprov.ProvEntry, []error, error) {
	log.Println("BuildIep")
	log.Println("reflink " + reflink)
	getSourceResources()
	generateResources()
	validateResources()
	xsdprov.ResrcJSON()
	xsdprov.ProvenanceRpt()
	return provreport, valerr, iepderr
}

func generateResources() (err error) {
	log.Println("Generate Resources")
	//spdx-license.xsd - Information Exchange Package XML Schema
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("spdx-license-iep.xsl", "spdx-ref.xsd", "spdx-license.xsd")
	check(err)
	//spdx-license-test-instance.xml - Information Exchange Package XML Instance
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("spdx-license-instance.xsl", "spdx-license.xsd", "spdx-license-test-instance.xml")
	check(err)
	//JSON
	//spdx-ref-xsd.json - JSON representation of spdx-ref.xsd
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("xsd-json.xsl", "spdx-ref.xsd", "spdx-ref-xsd.json")
	check(err)
	//spdx-license-iep-xsd.json - JSON representation of spdx-license-iep-xsd
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("xsd-json.xsl", "spdx-license.xsd", "spdx-license-iep-xsd.json")
	check(err)
	//spdx-license-test-instance.json - JSON representation license-test-instance.xml
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("xml-json.xsl", "spdx-license-test-instance.xml", "spdx-license-test-instance.json")
	check(err)
	//spdx-license-struct.go - Golang struct iep.go
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("go-gen-lic.xsl", "spdx-license.xsd", "spdx-license-struct.go")
	check(err)
	//spdx-license_test.go - Golang test iep.go
	provreport[time.Now().UnixNano()], err = xsdprov.GenerateResource("go-gen-lic-test.xsl", "spdx-license.xsd", "spdx-license_test.go")
	check(err)
	//Marshal instance
	provreport[time.Now().UnixNano()] = xsdprov.MarshalXML(temppath+resources["spdx-license-test-instance.xml"], resources["spdx-license-test-instance-golang.xml"], SPDXLicenseDatastruct)
	return err
}

func validateResources() (err error) {
	log.Println("Validate Resources")
	var errs []error
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-ref.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-license.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-license-test-instance.xml", "spdx-license.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-license-test-instance.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-license-test-instance-golang.xml", "spdx-license.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = xsdprov.ValidateFile("spdx-license-test-instance-golang.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	return err
}

func getSourceResources() {
	log.Println("getSourceResources")
	//Compare local copy of Ref XSD to Authoritative copy on GitHub
	var snr = "spdx-ref.xsd"
	log.Println(resources[snr])
	//tempfiles[snr] = temppath + resources[snr]
	//log.Println(tempfiles[snr])
	pe := xsdprov.LoadRemote(snr, temppath, reflink)
	provreport[time.Now().UnixNano()] = pe
	ped := xsdprov.CheckDigest(resources[snr], pe.Digest, tempdigests[snr])
	provreport[time.Now().UnixNano()] = ped
	if ped.Status == "Fail" {
		xsdprov.CopyFile(temppath+resources[snr], resources[snr])
		pcp := xsdprov.LoadRemote(snr, temppath, reflink)
		pcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = pcp
	}
	//Test Data
	var tdx = "spdx-test-data.xml"
	pex := xsdprov.LoadRemote(tdx, temppath, testlink)
	provreport[time.Now().UnixNano()] = pex
	pedx := xsdprov.CheckDigest(resources[tdx], pex.Digest, tempdigests[tdx])
	provreport[time.Now().UnixNano()] = pedx
	if pedx.Status == "Fail" {
		xsdprov.CopyFile(temppath+resources[tdx], resources[tdx])
		tcp := xsdprov.LoadRemote(tdx, temppath, testlink)
		tcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = tcp
	}
	tempfiles[tdx] = temppath + resources[tdx]
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
