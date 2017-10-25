package buntdb

import (
	"log"

	"github.com/tidwall/buntdb"
	"time"
)

var db *buntdb.DB

func init() {
	newClient()
}

func newClient() {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("data.db")
	//buntdb.Open(":memory:") // Open a file that does not persist to disk.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func GetValue(key string) string {

	var result string

	err := db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil{
			return err
		}
		result = val
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func SetValue(key string, value string, exp time.Duration) {

	opts := &buntdb.SetOptions{
		Expires: true,
		TTL: exp,
	}

	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, opts)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

}
