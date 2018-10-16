package xsdprov

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const pdb = "PROVDB"
const adb = "ADMIN"
const pd = "PROVDATA"

//DbSetup ...
func DbSetup(path string) (*bolt.DB, error) {
	log.Println("DbSetup, " + path)
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte(pdb))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(adb))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(pd))
		if err != nil {
			return fmt.Errorf("could not create days bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return db, nil
}
func updateDB(db *bolt.DB, bckt string, name string, data []byte) error {
	dberr := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt)).Put([]byte(name), data)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		log.Println("updateDB " + bckt + ", " + name + ", " + byteStr(data))
		return nil
	})
	fmt.Println("Updated Entry")
	return dberr
}
func queryDB(db *bolt.DB, bckt string, name string) (string, error) {
	log.Println("queryDB, " + bckt + ", " + name)
	var v = []byte{}
	verr := db.View(func(tx *bolt.Tx) error {
		val := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt)).Get([]byte(name))
		v = val
		fmt.Println(byteStr(v))
		return nil
	})
	fmt.Println("Found " + byteStr(v))
	return byteStr(v), verr
}
func jsonBytes(ifce interface{}, data string) ([]byte, error) {
	vs := []interface{}{ifce, data}
	encoded, err := json.Marshal(vs)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
func strBytes(data string) []byte {
	v := []byte(data)
	return v
}
func byteStr(data []byte) string {
	str := fmt.Sprintf("%s", data)
	return str
}
func checkDbE(e error) error {
	if e != nil {
		log.Println(err)
	}
	return e
}
