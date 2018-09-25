package xsdprov

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

<<<<<<< HEAD
var tempfiles = map[string]string{}
var resdigests = map[string]string{}
var tempdigests = map[string]string{}
var resources = map[string]string{}
var resourcedirs = map[string]string{}
var name string
var dbloc string
var cfg Cfg
var reflink string
var testlink string
var port string

//Tpath ...
var Tpath string

//Datastruct ...
var Datastruct interface{}

//Setup ...
func Setup(temppath string, resrces map[string]string, dirs map[string]string, dstruct interface{}) {
	cfg := getConfig()
	Datastruct = dstruct
	dbloc = "/tmp/" + cfg.Project
=======
var (
	resources    map[string]string
	resourcedirs map[string]string
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
	//Provreport ...
	Provreport = map[int64]ProvEntry{}
)

// InitXSDProv ...
func InitXSDProv(rsrcs map[string]string, rsrcdirs map[string]string, config string) {
	resources = rsrcs
	resourcedirs = rsrcdirs
	cfg := GetConfig(config)
	dbloc = cfg.Dbloc
	temppath = cfg.Temppath
>>>>>>> e6eb595232f7a1b0a8351ded210e2bbe11538545
	name = cfg.Project
	reflink = cfg.Reflink
	testlink = cfg.Testlink
	port = cfg.Port
<<<<<<< HEAD
	resources = resrces
	resourcedirs = dirs
	Tpath = temppath
	err = os.MkdirAll(dbloc+"/db", 0777)
=======
	resdigests = getDigests(resources, temppath, "Sha256")
	err := os.MkdirAll(dbloc, 0777)
>>>>>>> e6eb595232f7a1b0a8351ded210e2bbe11538545
	if err != nil {
		return
	}
	DirSetup()
<<<<<<< HEAD
	//BuildIep()
	//StartWeb(Tpath)
=======
	//CopyDirs(temppath, resourcedirs)
	db, err := DbSetup(dbloc + "/spdx-lic.db")
	check(err)
	// InitTempDir ...
	InitTempDir(db)
>>>>>>> e6eb595232f7a1b0a8351ded210e2bbe11538545
}

// InitTempDir ...
func InitTempDir(db *bolt.DB) (err error) {
	log.Println("TempDir ")
	tempdir, err := queryDB(db, "ADMIN", dbloc+"/spdx-lic.db")
	log.Println("TEMPDIR " + dbloc + "/spdx-lic.db")
	ferr := os.RemoveAll(dbloc + "/spdx-lic.db")
	err = ferr
<<<<<<< HEAD
	dberr := updateDB(db, "ADMIN", "tempdir", []byte(Tpath))
=======
	dberr := updateDB(db, "ADMIN", tempdir, []byte(dbloc+"/spdx-lic.db"))
>>>>>>> e6eb595232f7a1b0a8351ded210e2bbe11538545
	err = dberr
	return err
}

// DirSetup ...
func DirSetup() (e error) {
	log.Println("DirSetup")
	for _, rp := range resources {
<<<<<<< HEAD
		p := filepath.Dir(Tpath + rp)
		os.MkdirAll(p, os.ModePerm)
	}
	CopyDirs(Tpath, resourcedirs)
=======
		p := filepath.Dir(temppath + rp)
		os.MkdirAll(p, os.ModePerm)
	}
	CopyDirs(temppath, resourcedirs)
>>>>>>> e6eb595232f7a1b0a8351ded210e2bbe11538545
	return
}
