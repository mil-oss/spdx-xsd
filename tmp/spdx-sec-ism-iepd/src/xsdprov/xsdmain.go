package xsdprov

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	// Homeurl ...
	Homeurl     string
	tempfiles   = map[string]string{}
	resdigests  = map[string]string{}
	tempdigests = map[string]string{}
	name        string
	dbloc       string
	cfg         Cfg
	reflink     string
	testlink    string
	port        string
	// ProvDB ...
	ProvDB     *bolt.DB
	errorlist  []error
	provreport = map[int64]ProvEntry{}
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
	Homeurl = cfg.Homeurl
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
	InitTempDir()
}

// InitTempDir ...
func InitTempDir() (err error) {
	log.Println("InitTempDir")
	db, err := DbSetup(dbloc)
	check(err)
	var v = []byte{}
	verr := db.View(func(tx *bolt.Tx) error {
		val := tx.Bucket([]byte(pdb)).Bucket([]byte("ADMIN")).Get([]byte(temppath))
		v = val
		fmt.Println(byteStr(v))
		return nil
	})
	err = verr
	//td, err := queryDB("ADMIN", temppath)
	log.Println("TempPath " + string(v))
	//ferr := os.RemoveAll(dbloc)
	//err = ferr
	dberr := updateTransact("ADMIN", temppath, []byte(dbloc))
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
	GenerateResource(respath("iepxsdxsl"), respath("refxsd"), respath("iepxsd"))
	//test_instance.xml - Information Exchange Package XML Instance
	GenerateResource(respath("instancexsl"), respath("iepxsd"), respath("instancexml"))
	//JSON
	//iep.ref.json - JSON representation of ref.xsd
	GenerateResource(respath("xsdjsonxsl"), respath("refxsd"), respath("refxsdjson"))
	//iep.xsd.json - JSON representation of iep.xsd
	GenerateResource(respath("xsdjsonxsl"), respath("iepxsd"), respath("iepxsdjson"))
	//xml.json - JSON representation test_instance.xml
	GenerateResource(respath("xmljsonxsl"), respath("instancexml"), respath("instancejson"))
	//iep.xsd - Golang struct iep.go
	GenerateResource(respath("gogenxsdxsl"), respath("iepxsd"), respath("structgo"))
	//iep.xsd - Golang test iep.go
	GenerateResource(respath("gotestgenxsl"), respath("iepxsd"), respath("structtestgo"))
	//Marshal instance
	MarshalXML(respath("instancexml"), respath("instancegolangxml"), datastruct)

}

func respath(str string) string {
	return temppath + resources[str]
}

func validateResources() {
	log.Println("Validate Resources")
	ValidateFile("refxsd", "xmlschemaxsd")
	ValidateFile("iepxsd", "xmlschemaxsd")
	ValidateFile("instancexml", "iepxsd")
	ValidateFile("instancexml", "refxsd")
	ValidateFile("instancegolangxml", "iepxsd")
	ValidateFile("instancegolangxml", "refxsd")
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
