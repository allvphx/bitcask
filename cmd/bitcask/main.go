package main

import (
	"github.com/allvphx/bitcask"
	"log"
)

func main() {
	db, _ := bitcask.Open("/tmp/db")
	defer db.Close()
	db.Put([]byte("Hello"), []byte("World"))
	val, _ := db.Get([]byte("Hello"))
	log.Printf(string(val))
	//	Execute()
}
