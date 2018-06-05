package xsdprov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

var (
	testinstance string
	iepderr      error
	valerr       []error
	provreport   = map[int64]ProvEntry{}
)

//BuildIep ... Generate XML, Code and Test Artifacts
func BuildIep() (map[int64]ProvEntry, []error, error) {
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
	cerr := compress(Tpath, "/tmp/IEPD/"+name+".zip")
	check(cerr)
}

func provEntry(entrytype string, fpath string) ProvEntry {
	pe := ProvEntry{
		Timestamp: time.Now().UnixNano(),
		EntryType: entrytype,
		FilePath:  fpath,
	}
	return pe
}

func resrcJSON() []byte {
	rs, ferr := json.Marshal(resources)
	check(ferr)
	wferr := writeFile(resources["resources.json"], rs)
	check(wferr)
	rsferr := writeFile(Tpath+resources["resources.json"], rs)
	check(rsferr)
	return rs
}

func provenanceRpt() []byte {
	pr, err := json.Marshal(provreport)
	check(err)
	log.Println(Tpath + resources["provenance_report.json"])
	ferr := writeFile(Tpath+resources["provenance_report.json"], pr)
	check(ferr)
	return pr
}

func getSourceResources() {
	log.Println("getSourceResources")
	//Compare local copy of Ref XSD to Authoritative copy on GitHub
	var snr = "ref.xsd"
	tempfiles[snr] = Tpath + resources[snr]
	pe := loadRemote(snr, Tpath, reflink)
	provreport[time.Now().UnixNano()] = pe
	ped := checkDigest(resources[snr], pe.Digest, tempdigests[snr])
	provreport[time.Now().UnixNano()] = ped
	if ped.Status == "Fail" {
		CopyFile(Tpath+resources[snr], resources[snr])
		pcp := loadRemote(snr, Tpath, reflink)
		pcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = pcp
	}
	//Test Data
	var tdx = "test_data.xml"
	pex := loadRemote(tdx, Tpath, testlink)
	provreport[time.Now().UnixNano()] = pex
	pedx := checkDigest(resources[tdx], pex.Digest, tempdigests[tdx])
	provreport[time.Now().UnixNano()] = pedx
	if pedx.Status == "Fail" {
		CopyFile(Tpath+resources[tdx], resources[tdx])
		tcp := loadRemote(tdx, Tpath, testlink)
		tcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = tcp
	}
	tempfiles[tdx] = Tpath + resources[tdx]
}

func generateResources() {
	log.Println("Generate Resources")
	//GenerateResource - iep.xsd - Information Exchange Package XML Schema
	provreport[time.Now().UnixNano()], err = GenerateResource("iep_xsd.xsl", "ref.xsd", "iep.xsd")
	check(err)
	//test_instance.xml - Information Exchange Package XML Instance
	provreport[time.Now().UnixNano()], err = GenerateResource("xml_instance.xsl", "iep.xsd", "test_instance.xml")
	check(err)
	//JSON
	//iep.ref.json - JSON representation of ref.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource("xsd_json.xsl", "ref.xsd", "ref_xsd.json")
	check(err)
	//iep.xsd.json - JSON representation of iep.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource("xsd_json.xsl", "iep.xsd", "iep_xsd.json")
	check(err)
	//xml.json - JSON representation test_instance.xml
	provreport[time.Now().UnixNano()], err = GenerateResource("xml_json.xsl", "test_instance.xml", "test_instance.json")
	check(err)
	//iep.xsd - Golang struct iep.go
	provreport[time.Now().UnixNano()], err = GenerateResource("go-gen.xsl", "iep.xsd", "xsd-struct.go")
	check(err)
	//iep.xsd - Golang test iep.go
	provreport[time.Now().UnixNano()], err = GenerateResource("go-test-gen.xsl", "iep.xsd", "xsd-test.go")
	check(err)
	//Marshal instance
	provreport[time.Now().UnixNano()] = MarshalXML(Tpath+resources["test_instance.xml"], resources["test_instance-golang.xml"])
}

func validateResources() {
	log.Println("Validate Resources")
	var errs []error
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("ref.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("iep.xsd", "XMLSchema.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("test_instance.xml", "iep.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("test_instance.xml", "ref.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("test_instance-golang.xml", "iep.xsd")
	check(err)
	checka(errs)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("test_instance-golang.xml", "ref.xsd")
	check(err)
	checka(errs)
}

//GenerateResource ... generate IepXsd using XSLT
func GenerateResource(xslname string, xmlname string, resultname string) (ProvEntry, error) {
	log.Println("GenerateResource: " + resultname + "  XML Doc " + xmlname)
	pe := provEntry("GenerateResource", Tpath+resources[resultname])
	xslpath, xmlpath, resulTpath := getPaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := doTransform(xslpath, xmlpath)
	check(err)
	ferr := writeFile(resulTpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resulTpath, "Sha256"))
	tempfiles[resultname] = resulTpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe, err
}

//GenerateResourceParam ... generate IepXsd using XSLT
func GenerateResourceParam(xslname string, xmlname string, resultname string, testd string) ProvEntry {
	log.Println("GenerateResourceParam: " + resultname + "  XML Doc " + xmlname)
	pe := provEntry("GenerateResource", Tpath+resources[resultname])
	xslpath, xmlpath, resulTpath := getPaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := doTransformParam(xslpath, xmlpath, testd)
	check(err)
	ferr := writeFile(resulTpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resulTpath, "Sha256"))
	tempfiles[resultname] = resulTpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe
}

func getPaths(xslname string, xmlname string, resultname string) (string, string, string) {
	var xslpath = Tpath + resources[xslname]
	var xmlpath = Tpath + resources[xmlname]
	var resulTpath = Tpath + resources[resultname]
	if val, ok := tempfiles[xslname]; ok {
		xslpath = val
	}
	if val, ok := tempfiles[xmlname]; ok {
		xmlpath = val
	}
	log.Println("xslpath: " + xslpath)
	log.Println("xmlpath: " + xmlpath)
	log.Println("resulTpath: " + resulTpath)
	return xslpath, xmlpath, resulTpath
}

//MarshalXML ...
func MarshalXML(srcpath string, desTpath string) ProvEntry {
	var s = readStructXML(srcpath, Datastruct)
	var ft = filepath.Base(desTpath)
	tempfiles[ft] = Tpath + "/" + desTpath
	writeStructXML(tempfiles[ft], s)
	pe := provEntry("Marshal Data", tempfiles[ft])
	pe.Status = "Pass"
	pe.Digest = spaceMap(GetHash(tempfiles[ft], "Sha256"))
	return pe
}

//ValidateFile ... validate XML using XSD
func ValidateFile(xmlname string, xsdname string) (pe ProvEntry, errs []error, err error) {
	var xsdpath = Tpath + resources[xsdname]
	var xmlpath = Tpath + resources[xmlname]
	if val, ok := tempfiles[xsdname]; ok {
		xsdpath = val
	}
	if val, ok := tempfiles[xmlname]; ok {
		xmlpath = val
	}
	pe = provEntry("Validate", xmlpath)
	pe.XsdPath = xsdpath
	xmlstr, err := ioutil.ReadFile(xmlpath)
	vdata := ValidationData{XMLName: xmlname, XMLString: fmt.Sprintf("%s", xmlstr), XSDName: xsdname}
	valid, errs := ValidateXML(vdata)
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

func loadRemote(name string, path string, link string) ProvEntry {
	var refpath = path + resources[name]
	pe := provEntry("Load Remote Match", refpath)
	err = wgetFile(refpath, link)
	check(err)
	pe.Status = "Pass"
	pe.Message = "Loaded Remote Resource"
	tempdigests[name] = spaceMap(GetHash(refpath, "Sha256"))
	pe.Digest = tempdigests[name]
	return pe
}

func checkDigest(fpath string, auth string, test string) ProvEntry {
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
