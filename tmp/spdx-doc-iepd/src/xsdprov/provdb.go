package xsdprov

import (
	"encoding/binary"
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
	defer db.Close()
	return db, nil
}

func updateDB(bckt string, key string, data []byte) error {
	db, err := bolt.Open(dbloc, 0600, nil)
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	defer db.Close()
	dberr := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt)).Put([]byte(key), data)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		log.Println("updateDB " + bckt + ", " + key + ", " + byteStr(data))
		return nil
	})
	fmt.Println("Updated Entry")
	return dberr
}

func incrementDB(bckt string, data []byte) error {
	db, err := bolt.Open(dbloc, 0600, nil)
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	defer db.Close()
	dberr := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt))
		id, _ := b.NextSequence()
		err := b.Put(itob(int(id)), data)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		log.Println("updateDB " + bckt + ", " + string(id) + ", " + byteStr(data))
		return nil
	})
	fmt.Println("Updated Entry")
	return dberr
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func queryDB(bckt string, key string) (string, error) {
	db, err := bolt.Open(dbloc, 0600, nil)
	if err != nil {
		return "", fmt.Errorf("could not open db, %v", err)
	}
	log.Println("queryDB, " + bckt + ", " + key)
	var v = []byte{}
	verr := db.View(func(tx *bolt.Tx) error {
		val := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt)).Get([]byte(key))
		v = val
		fmt.Println(byteStr(v))
		return nil
	})
	fmt.Println("Found " + byteStr(v))
	return byteStr(v), verr
}

func updateTransact(bckt string, key string, data []byte) error {
	db, err := bolt.Open(dbloc, 0600, nil)
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	//defer db.Close()
	// Start a writable transaction.
	tx, err := db.Begin(true)
	check(err)
	defer tx.Rollback()
	// Use the transaction...
	terr := tx.Bucket([]byte(pdb)).Bucket([]byte(bckt)).Put([]byte(key), data)
	if terr != nil {
		return fmt.Errorf("could not insert entry: %v", terr)
	}
	log.Println("updateDB " + bckt + ", " + key + ", " + byteStr(data))
	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
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
