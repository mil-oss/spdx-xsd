package xsdprov

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	libxml2 "github.com/lestrrat/go-libxml2"
	xsd "github.com/lestrrat/go-libxml2/xsd"
)

// ReadStructXML ...
func ReadStructXML(filepath string, xsdstruct interface{}) interface{} {
	xf, ferr := ioutil.ReadFile(filepath)
	check(ferr)
	var strct = xsdstruct
	err := xml.Unmarshal([]byte(xf), &strct)
	check(err)
	return strct
}

// WriteStructXML ...
func WriteStructXML(filepath string, xsdstruct interface{}) string {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	var strct = xsdstruct
	output, err := xml.MarshalIndent(strct, "  ", "    ")
	check(err)
	var xmlrslt = []byte(xml.Header + string(output))
	ferr := ioutil.WriteFile(filepath, xmlrslt, 0666)
	check(ferr)
	return filepath
}

// Verify ... verify hash digest against known original
func Verify(verifydata VerifyData) bool {
	resdigests = getDigests(resources, temppath, "Sha256")
	//log.Println("Verify")
	log.Println("verifydata.ID " + verifydata.ID)
	log.Println("verifydata.Digest " + verifydata.Digest)
	log.Println("src.Digest " + resdigests[verifydata.ID])
	if resdigests[verifydata.ID] == verifydata.Digest {
		log.Println("Verification Successful")
		return true
	}
	return false
}

// ValidateXML ... validate XML against XSD
func ValidateXML(validationdata ValidationData) (bool, []error) {
	log.Println("ValidateXML")
	log.Println("xml: " + validationdata.XMLName)
	log.Println("xsd: " + validationdata.XSDName)
	var xsddoc, derr = xsd.ParseFromFile(temppath + resources[validationdata.XSDName])
	check(derr)
	//ioutil.ReadFile(validationdata.XMLPath)
	doc, err := libxml2.ParseString(validationdata.XMLString)
	check(err)
	if err := xsddoc.Validate(doc); err != nil {
		log.Println("Not Valid")
		return false, err.(xsd.SchemaValidationError).Errors()
	}
	return true, nil
}

// TransformXML ... generate a resource using XSLT
func TransformXML(transform TransformData) ([]byte, error) {
	log.Println("transformXML")
	log.Println("xml: " + transform.XMLName)
	log.Println("xmlpath: " + transform.XMLPath)
	log.Println("xsl: " + transform.XSLName)
	log.Println("xslpath: " + transform.XSLPath)
	cmd := exec.Cmd{
		Args: []string{"xsltproc", transform.XSLPath, transform.XMLPath},
		Env:  os.Environ(),
		Path: "xsltproc",
	}
	resultstring, err := cmd.Output()
	return resultstring, err
}

// DoTransform ...
func DoTransform(xslpath string, xmlpath string) ([]byte, error) {
	//log.Println("xslpath: " + xslpath)
	//log.Println("xmlpath: " + xmlpath)
	cmd := exec.Cmd{
		Args: []string{"xsltproc", xslpath, xmlpath},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}
	resultstring, err := cmd.Output()
	check(err)
	return resultstring, err
}

// DoTransformParam ...
func DoTransformParam(xslpath string, xmlpath string, testdata string) ([]byte, error) {
	//log.Println("xslpath: " + xslpath)
	//log.Println("xmlpath: " + xmlpath)
	cmd := exec.Cmd{
		Args: []string{"xsltproc", xslpath, xmlpath, "--stringparam", "TestData", testdata},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}
	resultstring, err := cmd.Output()
	check(err)
	return resultstring, err
}

func xmlOut(path string, data []byte) string {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err.Error()
	}
	return path
}
