package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

var world = []byte("world!")

func main() {
	db, err := bolt.Open("D:\\BattleField\\url-shortner\\bold.db", 0644, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "view.go : %v", err)
	}
	defer db.Close()
	key := []byte("hello")

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found !", world)
		}
		val := bucket.Get(key)
		fmt.Println(string(val))
		return nil
	})
}
