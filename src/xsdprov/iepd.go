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

func zipIEPD() {
	cerr := Compress(temppath, "/tmp/IEPD/"+name+".zip")
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

//ResrcJSON ...
func ResrcJSON() []byte {
	rs, ferr := json.Marshal(resources)
	check(ferr)
	wferr := WriteFile(resources["resources.json"], rs)
	check(wferr)
	rsferr := WriteFile(temppath+resources["resources.json"], rs)
	check(rsferr)
	return rs
}

//ProvenanceRpt ...
func ProvenanceRpt() []byte {
	pr, err := json.Marshal(provreport)
	check(err)
	log.Println(temppath + resources["provenance-report.json"])
	ferr := WriteFile(temppath+resources["provenance-report.json"], pr)
	check(ferr)
	return pr
}

//GenerateResource ... generate IepXsd using XSLT
func GenerateResource(xslname string, xmlname string, resultname string) (ProvEntry, error) {
	log.Println("GenerateResource: " + resultname + "  XML Doc " + xmlname)
	pe := provEntry("GenerateResource", temppath+resources[resultname])
	xslpath, xmlpath, resultemppath := getemppaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := DoTransform(xslpath, xmlpath)
	check(err)
	ferr := WriteFile(resultemppath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resultemppath, "Sha256"))
	tempfiles[resultname] = resultemppath
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
	pe := provEntry("GenerateResource", temppath+resources[resultname])
	xslpath, xmlpath, resultemppath := getemppaths(xslname, xmlname, resultname)
	pe.XslPath = xslpath
	doc, err := DoTransformParam(xslpath, xmlpath, testd)
	check(err)
	ferr := WriteFile(resultemppath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resultemppath, "Sha256"))
	tempfiles[resultname] = resultemppath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe
}

func getemppaths(xslname string, xmlname string, resultname string) (string, string, string) {
	var xslpath = temppath + resources[xslname]
	var xmlpath = temppath + resources[xmlname]
	var resultemppath = temppath + resources[resultname]
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

//MarshalXML ...
func MarshalXML(srcpath string, destemppath string, dstruct interface{}) ProvEntry {
	log.Println("MarshalXML: " + srcpath + "  to " + destemppath)
	var s = ReadStructXML(srcpath, dstruct)
	var ft = filepath.Base(destemppath)
	tempfiles[ft] = temppath + "/" + destemppath
	WriteStructXML(tempfiles[ft], s)
	pe := provEntry("Marshal Data", tempfiles[ft])
	pe.Status = "Pass"
	pe.Digest = spaceMap(GetHash(tempfiles[ft], "Sha256"))
	return pe
}

//ValidateFile ... validate XML using XSD
func ValidateFile(xmlname string, xsdname string) (pe ProvEntry, errs []error, err error) {
	var xsdpath = temppath + resources[xsdname]
	var xmlpath = temppath + resources[xmlname]
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
		return pe, errs, err
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

// LoadRemote ...
func LoadRemote(name string, path string, link string) ProvEntry {
	var refpath = path + resources[name]
	pe := provEntry("Load Remote Match", refpath)
	var err = WgetFile(refpath, link)
	check(err)
	pe.Status = "Pass"
	pe.Message = "Loaded Remote Resource"
	tempdigests[name] = spaceMap(GetHash(refpath, "Sha256"))
	pe.Digest = tempdigests[name]
	return pe
}

// CheckDigest ...
func CheckDigest(fpath string, auth string, test string) ProvEntry {
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
	return e
}

func checka(e []error) []error {
	if e != nil {
		fmt.Printf("error: %v\n", e)
	}
	return e
}
