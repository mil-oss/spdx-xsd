package xsdprov

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

//PORT ... Listen address
var PORT = 8080

var tempfiles = map[string]string{}
var resdigests = map[string]string{}
var tempdigests = map[string]string{}
var resources = map[string]string{}
var resourcedirs = map[string]string{}
var temppath string
var name string
var path string
var tpath string
var dbloc string
var cfg Cfg

//Datastruct ...
var Datastruct interface{}

//Setup ...
func Setup(pckgname string, assetpath string, resrces map[string]string, dirs map[string]string, dstruct interface{}) {
	cfg := getConfig()
	Datastruct = dstruct
	dbloc = "/tmp/" + pckgname
	temppath = "/tmp/IEPD/iepd"
	path = assetpath
	name = pckgname
	resources = resrces
	resourcedirs = dirs
	tpath = temppath + "/" + path
	err = os.MkdirAll(dbloc+"/db", 0777)
	if err != nil {
		return
	}
	db, err := DbSetup(dbloc)
	check(err)
	e := TempDir(db)
	check(e)
	DirSetup()
	BuildIep()
	StartWeb(path, temppath)
}

//TempDir ...
func TempDir(db *bolt.DB) (err error) {
	log.Println("TempDir ")
	tempdir, err := queryDB(db, "ADMIN", "tempdir")
	log.Println("TEMPDIR " + tempdir)
	ferr := os.RemoveAll(tempdir)
	err = ferr
	dberr := updateDB(db, "ADMIN", "tempdir", []byte(temppath))
	err = dberr
	return err
}

// DirSetup ...
func DirSetup() (e error) {
	for _, rp := range resources {
		p := filepath.Dir(tpath + rp)
		os.MkdirAll(p, os.ModePerm)
	}
	CopyDirs(path, tpath, resourcedirs)
	return
}
