package xsdprov

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var (
	datastruct   interface{}
	rsrcs        []Resource
	rsrcdirs     []Resource
	resources    = map[string]string{}
	resourcedirs = map[string]string{}
	sources      = map[string]string{}
	tempdir      string
	temppath     string
	tempfiles    = map[string]string{}
	resdigests   = map[string]string{}
	tempdigests  = map[string]string{}
	name         string
	dbloc        string
	cfg          Cfg
	reflink      string
	testlink     string
	port         string
	db           *bolt.DB
	errorlist    []error
	provreport   = map[int64]ProvEntry{}
)

// InitXSDProv ...
func InitXSDProv(config string) {
	cfg := GetConfig(config)
	for r := range cfg.Resources {
		resources[cfg.Resources[r].Name] = cfg.Resources[r].Path
		sources[cfg.Resources[r].Name] = cfg.Resources[r].Src
	}
	for r := range cfg.Directories {
		resourcedirs[cfg.Directories[r].Name] = cfg.Directories[r].Path
		resourcedirs[cfg.Directories[r].Name] = cfg.Directories[r].Src
	}
	dbloc = cfg.Dbloc
	tempdir = cfg.Tempdir
	temppath = cfg.Temppath
	name = cfg.Project
	reflink = cfg.Reflink
	testlink = cfg.Testlink
	port = cfg.Port
	resdigests = getDigests(resources, temppath, "Sha256")
	dbp := filepath.Dir(dbloc)
	err := os.MkdirAll(dbp, 0777)
	if err != nil {
		return
	}
	DirSetup()
	db, err := DbSetup(cfg.Dbloc)
	check(err)
	// InitTempDir ...
	InitTempDir(db)
}

// InitTempDir ...
func InitTempDir(db *bolt.DB) (err error) {
	log.Println("InitTempDir")
	tempdir, err := queryDB(db, "ADMIN", dbloc)
	log.Println("TEMPDIR " + dbloc)
	ferr := os.RemoveAll(dbloc)
	err = ferr
	dberr := updateDB(db, "ADMIN", tempdir, []byte(dbloc))
	err = dberr
	return err
}

// DirSetup ...
func DirSetup() (e error) {
	log.Println("DirSetup")
	for f := range resources {
		dest := filepath.Dir(temppath + resources[f])
		os.MkdirAll(dest, os.ModePerm)
		if sources[f] != "" {
			CopyFile(sources[f], temppath+resources[f])
		}
	}
	CopyDirs(temppath, resourcedirs)
	return
}

//BuildIep ... Generate XML, Code and Test Artifacts
func BuildIep(dstruct interface{}) (map[int64]ProvEntry, []error) {
	datastruct = dstruct
	GetSourceResources()
	generateResources()
	validateResources()
	ResrcJSON(respath("resourcesjson"))
	ProvenanceRpt()
	ZipIEPD(temppath, "tmp/"+resources["zipiepd"])
	return provreport, errorlist
}

func generateResources() {
	log.Println("Generate Resources")
	//GenerateResource - iep.xsd - Information Exchange Package XML Schema
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("iepxsdxsl"), respath("refxsd"), respath("iepxsd"))
	check(err)
	//test_instance.xml - Information Exchange Package XML Instance
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("instancexsl"), respath("iepxsd"), respath("instancexml"))
	check(err)
	//JSON
	//iep.ref.json - JSON representation of ref.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("xsdjsonxsl"), respath("refxsd"), respath("refxsdjson"))
	check(err)
	//iep.xsd.json - JSON representation of iep.xsd
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("xsdjsonxsl"), respath("iepxsd"), respath("iepxsdjson"))
	check(err)
	//xml.json - JSON representation test_instance.xml
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("xmljsonxsl"), respath("instancexml"), respath("instancejson"))
	check(err)
	//iep.xsd - Golang struct iep.go
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("gogenxsdxsl"), respath("iepxsd"), respath("structgo"))
	check(err)
	//iep.xsd - Golang test iep.go
	provreport[time.Now().UnixNano()], err = GenerateResource(respath("gotestgenxsl"), respath("iepxsd"), respath("structtestgo"))
	check(err)
	//Marshal instance
	provreport[time.Now().UnixNano()] = MarshalXML(respath("instancexml"), respath("instancegolangxml"), datastruct)
}

func respath(str string) string {
	return temppath + resources[str]
}

func validateResources() {
	log.Println("Validate Resources")
	var errs []error
	var err error
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("refxsd", "xmlschemaxsd")
	valerr(errs, err)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("iepxsd", "xmlschemaxsd")
	valerr(errs, err)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("instancexml", "iepxsd")
	valerr(errs, err)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("instancexml", "refxsd")
	valerr(errs, err)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("instancegolangxml", "iepxsd")
	valerr(errs, err)
	provreport[time.Now().UnixNano()], errs, err = ValidateFile("instancegolangxml", "refxsd")
	valerr(errs, err)
}
func valerr(er []error, e error) {
	if er != nil {
		fmt.Printf("error: %v\n", e)
		for ve := range er {
			errorlist = append(errorlist, er[ve])
		}
	}
	if e != nil {
		fmt.Printf("error: %v\n", e)
		errorlist = append(errorlist, e)
	}
}

func beforeStr(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
