package shorturl

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

// Objective:
// Create three funcs main(), FindUrl(), CreateUrl()

// main ... is used to create the entries manually and test the UpdateURL and FindURL methods.
func main() {
	src := "1675a480"
	// dest := "https://golang.org/pkg/net/http"

	// err := CreateURL(src, dest)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error in CreateURL() :%v", err)
	// }
	// fmt.Println("src-dest pair has been successfully updated")

	dest, found, err := FindURL(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in FindURL() :%v", err)
		os.Exit(1)
	}
	if found {
		fmt.Printf("%s --> %s\n ", src, dest)
	} else {
		fmt.Println("Not found !")
	}
}

// CreateURL ... will insert a (src,dest) in Db, for now it will be called from main.
func CreateURL(src, dest string) error {
	db, err := bolt.Open("C:\\Users\\Dictator\\Desktop\\BattleField\\go-shorturl\\src\\database\\shorturl.db", 0644, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening the Db : %v", err)
		return err
	}
	defer db.Close()
	// Store src and destination.
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("srcdest"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(src), []byte(dest))
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// FindURL ... will be called with short string and it will query the database and will return the database with dest and ok
func FindURL(src string) (string, bool, error) {
	db, err := bolt.Open("C:\\Users\\Dictator\\Desktop\\BattleField\\go-shorturl\\src\\database\\shorturl.db", 0644, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while maintaining the connection with the db: %v", err)
		return "", false, err
	}
	defer db.Close()
	var dest []byte
	// Retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("srcdest"))
		if bucket == nil {
			return fmt.Errorf("Bucket srcdest not found! %v", bucket)
		}
		dest = bucket.Get([]byte(src))
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while getting dest %v", err)
		return string(dest), false, err
	}
	if dest == nil {
		return string(dest), false, nil
	}
	return string(dest), true, nil
}
