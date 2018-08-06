package xsdprov

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

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
	name = cfg.Project
	reflink = cfg.Reflink
	testlink = cfg.Testlink
	port = cfg.Port
	resdigests = getDigests(resources, temppath, "Sha256")
	err := os.MkdirAll(dbloc, 0777)
	if err != nil {
		return
	}
	DirSetup()
	CopyDirs(temppath, resourcedirs)
	db, err := DbSetup(dbloc + "/spdx-lic.db")
	check(err)
	// InitTempDir ...
	InitTempDir(db)
}

// InitTempDir ...
func InitTempDir(db *bolt.DB) (err error) {
	log.Println("TempDir ")
	tempdir, err := queryDB(db, "ADMIN", dbloc+"/spdx-lic.db")
	log.Println("TEMPDIR " + dbloc + "/spdx-lic.db")
	ferr := os.RemoveAll(dbloc + "/spdx-lic.db")
	err = ferr
	dberr := updateDB(db, "ADMIN", tempdir, []byte(dbloc+"/spdx-lic.db"))
	err = dberr
	return err
}

// DirSetup ...
func DirSetup() (e error) {
	log.Println("DirSetup")
	for _, rp := range resources {
		p := filepath.Dir(temppath + rp)
		os.MkdirAll(p, os.ModePerm)
	}
	//CopyDirs(temppath, resourcedirs)
	return
}
