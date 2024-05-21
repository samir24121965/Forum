package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

// Creating Key/Value template
var Tmpls = map[string]string{
	"base":     "./templates/base.html",
	"nav":      "./templates/nav.html",
	"login":    "./templates/login.html",
	"register": "./templates/register.html",
}

// return the template
func renderTemplate(w http.ResponseWriter, files []string, data interface{}) {
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", data)
}

// Find the user in the database
func (db *database) findUser(name string, email string) (bool, error) {
	var count uint64
	stmt, err := db.db.Prepare(`
	select count(id) as count 
	from 
		user 
	where 
			name = $1
		or email = $2
	`)
	if err != nil {
		return false, err
	}
	row := stmt.QueryRow(name, email)
	row.Scan(&count)
	fmt.Println(name, email, count)

	return count >= 1, nil
}

// Find the user by login
func (db *database) findUserLogin(name string, password string) (bool, error) {
	var count uint64
	stmt, err := db.db.Prepare(`
	select count(id) as count 
	from 
		user 
	where 
			name = $1
		and password = $2
	`)
	if err != nil {
		return false, err
	}
	row := stmt.QueryRow(name, password)
	row.Scan(&count)
	fmt.Println(name, password, count)

	return count == 1, nil
}

// get user from the database with ID
func (db *database) getUserByID(id uint64) (*User, error) {
	stmt, err := db.db.Prepare(`
	select id, name, email, password 
	from 
		user 
	where 
			id = $1
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	u := User{}
	err = row.Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Create the user in the database
func (db *database) createUser(u *User) error {
	insert := `
	insert into 
	user 
		(name, email, password, created_at) 
	values 
		(?, ?, ?, ?)
	`
	userID, err := db.db.Exec(insert, u.Name, strings.ToLower(u.Email), u.Password, time.Now().Format("2006-01-02 15:04"))
	if err != nil {
		return err
	}
	u.Id, err = userID.LastInsertId()
	// fmt.Println(u.Id)
	return err
}
