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

// GetSourceResources ...
func GetSourceResources() {
	log.Println("getSourceResources")
	//Compare local copy of Ref XSD to Authoritative copy on GitHub
	var snr = "refxsd"
	log.Println(resources[snr])
	tempfiles[snr] = temppath + resources[snr]
	log.Println(tempfiles[snr])
	pe := LoadRemote(snr, tempfiles[snr], reflink)
	provreport[time.Now().UnixNano()] = pe
	ped := CheckDigest(resources[snr], pe.Digest, tempdigests[snr])
	provreport[time.Now().UnixNano()] = ped
	if ped.Status == "Fail" {
		CopyFile(temppath+resources[snr], resources[snr])
		pcp := LoadRemote(snr, tempfiles[snr], reflink)
		pcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = pcp
	}
	//Test Data
	var tdx = "testdataxml"
	tempfiles[tdx] = temppath + resources[tdx]
	pex := LoadRemote(tdx, tempfiles[tdx], testlink)
	provreport[time.Now().UnixNano()] = pex
	pedx := CheckDigest(resources[tdx], pex.Digest, tempdigests[tdx])
	provreport[time.Now().UnixNano()] = pedx
	if pedx.Status == "Fail" {
		CopyFile(temppath+resources[tdx], resources[tdx])
		tcp := LoadRemote(tdx, tempfiles[tdx], testlink)
		tcp.Message = "Resource Updated"
		provreport[time.Now().UnixNano()] = tcp
	}
	tempfiles[tdx] = temppath + resources[tdx]
}

// ZipIEPD ...
func ZipIEPD(dirpath string, path string) {
	cerr := Compress(dirpath, path)
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
func ResrcJSON(path string) []byte {
	rs, ferr := json.Marshal(resources)
	check(ferr)
	rsferr := WriteFile(path, rs)
	check(rsferr)
	return rs
}

//ProvenanceRpt ...
func ProvenanceRpt() []byte {
	log.Println("ProvenanceRpt " + temppath + resources["provenancereportjson"])
	pr, err := json.Marshal(provreport)
	check(err)
	ferr := WriteFile(temppath+resources["provenancereportjson"], pr)
	check(ferr)
	return pr
}

//GenerateResource ... generate IepXsd using XSLT
func GenerateResource(xslpath string, xmlpath string, resultpath string) (ProvEntry, error) {
	log.Println("GenerateResource: " + resultpath + "  from " + xmlpath + "  with " + xslpath)
	var resultname = filepath.Base(resultpath)
	pe := provEntry("GenerateResource", resultname)
	pe.XslPath = xslpath
	doc, err := DoTransform(xslpath, xmlpath)
	check(err)
	ferr := WriteFile(resultpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resultpath, "Sha256"))
	tempfiles[resultname] = resultpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe, err
}

//GenerateResourceParam ... generate IepXsd using XSLT
func GenerateResourceParam(xslpath string, xmlpath string, resultpath string, paramstr string) ProvEntry {
	log.Println("GenerateResourceParam: " + resultpath + "  from " + xmlpath)
	var resultname = filepath.Base(resultpath)
	pe := provEntry("GenerateResource", resultname)
	pe.XslPath = xslpath
	doc, err := DoTransformParam(xslpath, xmlpath, paramstr)
	check(err)
	ferr := WriteFile(resultpath, doc)
	check(ferr)
	tempdigests[resultname] = spaceMap(GetHash(resultpath, "Sha256"))
	tempfiles[resultname] = resultpath
	pe.Digest = tempdigests[resultname]
	pe.Status = "Pass"
	if err != nil {
		pe.Status = "Fail"
	}
	return pe
}

//MarshalXML ...
func MarshalXML(srcpath string, destpath string, dstruct interface{}) ProvEntry {
	log.Println("MarshalXML: " + srcpath + "  to " + destpath)
	var s = ReadStructXML(srcpath, dstruct)
	var name = filepath.Base(destpath)
	WriteStructXML(destpath, s)
	pe := provEntry("Marshal Data", name)
	pe.Status = "Pass"
	pe.Digest = spaceMap(GetHash(destpath, "Sha256"))
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
	pe := provEntry("Load Remote Match", path)
	var err = WgetFile(path, link)
	check(err)
	pe.Status = "Pass"
	pe.Message = "Loaded Remote Resource"
	tempdigests[name] = spaceMap(GetHash(path, "Sha256"))
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
