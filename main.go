package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// le underscore (_) c'est un alias
)

type config struct {
	port string
	db   database
}

var cfg config

func main() {

	cfg.port = "8888"
	cfg.db.dbName = "forum.db"
	cfg.db.driver = "sqlite3"
	err := cfg.db.openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer cfg.db.db.Close()

	err = createDB(cfg.db.db)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Good")
	// found, err := cfg.db.findUser("samir", "samir")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(found)
	// u, err := cfg.db.getUserByID(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(u)

	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates", fs))

	router()

	fmt.Println("http://localhost:" + cfg.port)
	if err := http.ListenAndServe(":"+cfg.port, nil); err != nil {
		log.Fatal(err)
	}

}
