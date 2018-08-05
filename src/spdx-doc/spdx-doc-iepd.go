package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
	"xsdprov"
)

var (
	dbloc        string
	temppath     string
	name         string
	reflink      string
	testlink     string
	port         string
	tpath        string
	resources    map[string]string
	resourcedirs map[string]string
	tempfiles    map[string]string
	tempdigests  map[string]string
)

//SPDXDocDatastruct ...
var SPDXDocDatastruct interface{}

//SetupSPDXDocument ...
func SetupSPDXDocument(resrces map[string]string, dirs map[string]string) {
	cfg := xsdprov.GetConfig("config/spdx-doc-cfg.json")
	dbloc = "/tmp/" + cfg.Project
	temppath = "/tmp/IEPD/iepd"
	name = cfg.Project
	reflink = cfg.Reflink
	testlink = cfg.Testlink
	port = cfg.Port
	resources = resrces
	resourcedirs = dirs
	tpath = temppath + "/"
	var err = os.MkdirAll(dbloc+"/db", 0777)
	if err != nil {
		return
	}
	db, err := xsdprov.DbSetup(dbloc)
	check(err)
	e := xsdprov.TempDir(db)
	check(e)
	xsdprov.DirSetup()
	BuildDocumentIep()
	xsdprov.StartWeb(temppath)
}

//BuildDocumentIep ... Generate XML, Code and Test Artifacts
func BuildDocumentIep() (map[int64]xsdprov.ProvEntry, []error, error) {
	log.Println("BuildIep")
	log.Println("reflink " + reflink)
	getSourceResources()
	generateResources()
	validateResources()
	resrcJSON()
	provenanceRpt()
	zipIEPD()
	return provreport, valerr, iepderr
}

func zipIEPD() {
	cerr := xsdprov.Compress(tpath, "/tmp/IEPD/"+name+".zip")
	check(cerr)
}

func provEntry(entrytype string, fpath string) xsdprov.ProvEntry {
	pe := xsdprov.ProvEntry{
		Timestamp: time.Now().UnixNano(),
		EntryType: entrytype,
		FilePath:  fpath,
	}
	return pe
}

func resrcJSON() []byte {
	rs, ferr := json.Marshal(resources)
	check(ferr)
	wferr := xsdprov.WriteFile(resources["resources.json"], rs)
	check(wferr)
	rsferr := xsdprov.WriteFile(tpath+resources["resources.json"], rs)
	check(rsferr)
	return rs
}

func provenanceRpt() []byte {
	pr, err := json.Marshal(provreport)
	check(err)
	log.Println(tpath + resources["provenance-report.json"])
	ferr := xsdprov.WriteFile(tpath+resources["provenance-report.json"], pr)
	check(ferr)
	return pr
}

func getSourceResources() {
	log.Println("getSourceResources")
	//Compare local copy of Ref XSD to Authoritative copy on GitHub
	var snr = "spdx-ref.xsd"
	tempfiles[snr] = tpath + resources[snr]
	pe := loadRemote(snr, tpath, reflink)
	provreport[time.Now().UnixNano()] = pe
	ped := checkDigest(resources[snr], pe.Digest, tempdigests[snr])
	provreport[time.Now().UnixNano()] = ped
	if ped.Status == "Fail" {
		xsdprov.CopyFile(tpath+resources[snr], resources[snr])
		pcp := loadRemote(snr, tpath, reflink)
		pcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = pcp
	}
	//Test Data
	var tdx = "spdx-test-data.xml"
	pex := loadRemote(tdx, tpath, testlink)
	provreport[time.Now().UnixNano()] = pex
	pedx := checkDigest(resources[tdx], pex.Digest, tempdigests[tdx])
	provreport[time.Now().UnixNano()] = pedx
	if pedx.Status == "Fail" {
		xsdprov.CopyFile(tpath+resources[tdx], resources[tdx])
		tcp := loadRemote(tdx, tpath, testlink)
		tcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = tcp
	}
	tempfiles[tdx] = tpath + resources[tdx]
}

func generateResources() (err error) {
	log.Println("Generate Resources")
	//GenerateResource - spdx-doc.xsd - Information Exchange Package XML Schema
	provreport[time.Now().UnixNano()], err = GenerateResource("spdx-doc-iep.xsl", "spdx-ref.xsd", "spdx-doc.xsd")
	check(err)
	//doc-test-instance.xml - Information Exchange Package XML Instance
	provreport[time.Now().UnixNano()], err = GenerateResource("spdx-doc-instance.xsl", "spdx-doc.xsd", "spdx-doc-test-instance.xml")
	check(err)
	//JSON
	//spdx-ref-xsd.json - JSON representation of ref.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource("xsd-json.xsl", "spdx-ref.xsd", "spdx-ref-xsd.json")
	check(err)
	//spdx-doc-iep-xsd.json - JSON representation of spdx-doc-iep.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource("xsd-json.xsl", "spdx-doc.xsd", "spdx-doc-iep-xsd.json")
	check(err)
	//spdx-doc-test-instance.json - JSON representation doc-test-instance.xml
	provreport[time.Now().UnixNano()], err = GenerateResource("xml-json.xsl", "spdx-doc-test-instance.xml", "spdx-doc-test-instance.json")
	check(err)
	//spdx-doc-struct.go - Golang struct
	provreport[time.Now().UnixNano()], err = GenerateResource("go-gen-doc.xsl", "spdx-doc.xsd", "spdx-doc-struct.go")
	check(err)
	//spdx-doc_test.go - Golang test
	provreport[time.Now().UnixNano()], err = GenerateResource("go-gen-doc-test.xsl", "spdx-doc.xsd", "spdx-doc_test.go")
	check(err)
	//Marshal instance
	provreport[time.Now().UnixNano()] = MarshalXML(tpath+resources["spdx-doc-test-instance.xml"], resources["spdx-doc-test-instance-golang.xml"], SPDXDocDatastruct)
	return err
}

func validateResources() (errs []error, err error) {
	log.Println("Validate Resources")
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-ref.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-doc.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-doc-test-instance.xml", "spdx-doc.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-doc-test-instance.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-doc-test-instance-golang.xml", "spdx-doc.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("spdx-doc-test-instance-golang.xml", "spdx-ref.xsd")
	check(err)
	checka(errs)
	return errs, err
}

//GenerateResource ... generate IepXsd using XSLT
func GenerateResource(xslname string, xmlname string, resultname string) (xsdprov.ProvEntry, error) {
	log.Println("GenerateResource: " + resultname + "  XML Doc " + xmlname)
	pe := provEntry("GenerateResource", tpath+resources[resultname])
	xslpath, xmlpath, resultpath := getPaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := xsdprov.DoTransform(xslpath, xmlpath)
	check(err)
	ferr := xsdprov.WriteFile(resultpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(xsdprov.GetHash(resultpath, "Sha256"))
	tempfiles[resultname] = resultpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe, err
}

//GenerateResourceParam ... generate IepXsd using XSLT
func GenerateResourceParam(xslname string, xmlname string, resultname string, testd string) xsdprov.ProvEntry {
	log.Println("GenerateResourceParam: " + resultname + "  XML Doc " + xmlname)
	pe := provEntry("GenerateResource", tpath+resources[resultname])
	xslpath, xmlpath, resultpath := getPaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := xsdprov.DoTransformParam(xslpath, xmlpath, testd)
	check(err)
	ferr := xsdprov.WriteFile(resultpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(xsdprov.GetHash(resultpath, "Sha256"))
	tempfiles[resultname] = resultpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe
}

func getPaths(xslname string, xmlname string, resultname string) (string, string, string) {
	var xslpath = tpath + resources[xslname]
	var xmlpath = tpath + resources[xmlname]
	var resultpath = tpath + resources[resultname]
	if val, ok := tempfiles[xslname]; ok {
		xslpath = val
	}
	if val, ok := tempfiles[xmlname]; ok {
		xmlpath = val
	}
	log.Println("xslpath: " + xslpath)
	log.Println("xmlpath: " + xmlpath)
	log.Println("resultpath: " + resultpath)
	return xslpath, xmlpath, resultpath
}

//MarshalXML ...
func MarshalXML(srcpath string, destpath string, dstruct interface{}) xsdprov.ProvEntry {
	log.Println("MarshalXML: " + srcpath + "  to " + destpath)
	var s = xsdprov.ReadStructXML(srcpath, dstruct)
	var ft = filepath.Base(destpath)
	tempfiles[ft] = tpath + "/" + destpath
	xsdprov.WriteStructXML(tempfiles[ft], s)
	pe := provEntry("Marshal Data", tempfiles[ft])
	pe.Status = "Pass"
	pe.Digest = spaceMap(xsdprov.GetHash(tempfiles[ft], "Sha256"))
	return pe
}

//ValidateFile ... validate XML using XSD
func ValidateFile(xmlname string, xsdname string) (pe xsdprov.ProvEntry, errs []error, err error) {
	var xsdpath = tpath + resources[xsdname]
	var xmlpath = tpath + resources[xmlname]
	if val, ok := tempfiles[xsdname]; ok {
		xsdpath = val
	}
	if val, ok := tempfiles[xmlname]; ok {
		xmlpath = val
	}
	pe = provEntry("Validate", xmlpath)
	pe.XsdPath = xsdpath
	xmlstr, err := ioutil.ReadFile(xmlpath)
	vdata := xsdprov.ValidationData{XMLName: xmlname, XMLString: fmt.Sprintf("%s", xmlstr), XSDName: xsdname}
	valid, errs := xsdprov.ValidateXML(vdata)
	checka(errs)
	if !valid {
		pe.Status = "Fail"
		pe.Valid = false
		pe.Errors = jsonList(errs)
		return pe, valerr, iepderr
	}
	pe.Valid = true
	pe.Status = "Pass"
	return pe, nil, nil
}

func jsonList(errs []error) []string {
	errlist := []string{}
	for _, errorMessage := range errs {
		errlist = append(errlist, errorMessage.Error())
	}
	return errlist
}

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func loadRemote(name string, path string, link string) xsdprov.ProvEntry {
	var refpath = path + resources[name]
	pe := provEntry("Load Remote Match", refpath)
	var err = xsdprov.WgetFile(refpath, link)
	check(err)
	pe.Status = "Pass"
	pe.Message = "Loaded Remote Resource"
	tempdigests[name] = spaceMap(xsdprov.GetHash(refpath, "Sha256"))
	pe.Digest = tempdigests[name]
	return pe
}

func checkDigest(fpath string, auth string, test string) xsdprov.ProvEntry {
	pe := provEntry("Authenticity Check", fpath)
	pe.Status = "Pass"
	pe.Digest = test
	pe.Message = filepath.Base(fpath) + " matches authoritative source"
	if auth != test {
		pe.Status = "Fail"
		pe.Message = filepath.Base(fpath) + " does NOT match authoritative source"
		return pe
	}
	return pe
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
