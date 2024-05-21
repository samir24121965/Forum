package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
// )

// type User struct {
// 	Id         int
// 	Name       string
// 	Created_at string
// }

// func main() {
// 	// var forumDB *sql.DB
// 	forumDB, err := sql.Open("sqlite3", "forum.db") // Specification of driver (sqlite3) and name of database (forum.db)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer forumDB.Close()

// 	err = forumDB.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := forumDB.Query(`
// 	SELECT 
// 		id, name, strftime('%d/%m/%Y', created_at) as created_at 
// 	FROM 
// 		user 
// 	where 
// 		date(created_at) between date('2023-03-28') and date('2023-03-29')
// 	`)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	user := User{}
// 	for rows.Next() {
// 		rows.Scan(&user.Id, &user.Name, &user.Created_at)
// 		fmt.Printf("User: %v\n", user)
// 	}
// }
