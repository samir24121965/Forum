package main

import (
	"database/sql"
	"fmt"
)

type HandleDB interface {
	openDB() (*sql.DB, error)
}

type database struct {
	db     *sql.DB
	dbName string
	driver string
}

func (db *database) openDB() error {
	dbase, err := sql.Open(db.driver, db.dbName) // Specification of driver (sqlite3) and name of database (forum.db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = dbase.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}

	db.db = dbase
	return nil
}

func createDB(db *sql.DB) error {

	// COLLATE NOCASE : case insensitive
	var schema = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE COLLATE NOCASE,
		email TEXT NOT NULL UNIQUE COLLATE NOCASE,
		password TEXT NOT NULL,
		is_disabled INTEGER NOT NULL DEFAULT 0,
		is_deleted INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL,
		updated_at TEXT DEFAULT NULL
	);

	CREATE TABLE IF NOT EXISTS post (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		nb_likes INTEGER NOT NULL DEFAULT 0, -- To avoid SQL COUNT each time
		nb_dislikes INTEGER NOT NULL DEFAULT 0, -- To avoid SQL COUNT each time
		created_at TEXT NOT NULL,
		updated_at TEXT DEFAULT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id)
	);

	CREATE TABLE IF NOT EXISTS category (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		nb_posts INTEGER NOT NULL DEFAULT 0, -- To avoid SQL COUNT each time
		created_at TEXT NOT NULL,
		updated_at TEXT DEFAULT NULL
	);

	CREATE TABLE IF NOT EXISTS post_category (
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		PRIMARY KEY (post_id, category_id),
		FOREIGN KEY(post_id) REFERENCES post(id),
		FOREIGN KEY(category_id) REFERENCES category(id)
	);

	CREATE TABLE IF NOT EXISTS comment (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		parent_id INTEGER NULL,
		content TEXT NOT NULL,
		nb_likes INTEGER NOT NULL DEFAULT 0, -- To avoid SQL COUNT each time
		nb_dislikes INTEGER NOT NULL DEFAULT 0, -- To avoid SQL COUNT each time
		is_deleted INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL,
		updated_at TEXT DEFAULT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id),
		FOREIGN KEY(parent_id) REFERENCES comment(id)
	);

	CREATE TABLE IF NOT EXISTS post_reaction (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		is_liked INTEGER NOT NULL,
		created_at TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
	);

	CREATE TABLE IF NOT EXISTS comment_reaction (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		is_liked INTEGER NOT NULL,
		created_at TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(comment_id) REFERENCES comment(id)
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
