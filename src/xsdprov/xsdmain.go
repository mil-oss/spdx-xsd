package xsdprov

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

var tempfiles = map[string]string{}
var resdigests = map[string]string{}
var tempdigests = map[string]string{}
var resources = map[string]string{}
var resourcedirs = map[string]string{}
var temppath string
var name string
var tpath string
var dbloc string
var cfg Cfg
var reflink string
var testlink string
var port string

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
	log.Println("DirSetup")
	for _, rp := range resources {
		p := filepath.Dir(tpath + rp)
		os.MkdirAll(p, os.ModePerm)
	}
	CopyDirs(tpath, resourcedirs)
	return
}
