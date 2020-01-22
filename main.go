package BC

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

var world = []byte("world!")

func avc() {
	db, err := bolt.Open("D:\\BattleField\\url-shortner\\bold.db", 0644, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "db.go : %v", err)
		os.Exit(1)
	}
	defer db.Close()

	key := []byte("hello")
	value := []byte("Hello world!")

	//store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v", err)
	}

	// retreive the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found !", world)
		}
		val := bucket.Get(key)
		fmt.Println(string(val))
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v", err)
	}
}
