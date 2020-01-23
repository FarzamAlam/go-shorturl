package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

var world = []byte("world!")

func main() {
	db, err := bolt.Open("C:\\Users\\Dictator\\Desktop\\BattleField\\url-shortner\\bold.db", 0644, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "view.go : %v", err)
	}
	defer db.Close()

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found !", world)
		}
		val := bucket.Get([]byte("farz"))
		fmt.Println(string(val))
		return nil
	})
}
