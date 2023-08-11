package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// TODO: Do validation

var (
	db *sql.DB
)

func init() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp", // is tcp by default
		Addr:                 "localhost:3306",
		DBName:               "go_database",
		AllowNativePasswords: true,
	}

	var openErr error
	db, openErr = sql.Open("mysql", cfg.FormatDSN())

	if IsNotNil(openErr) {
		FLog(openErr)
	}

	if pingErr := db.Ping(); IsNotNil(pingErr) {
		FLog(pingErr)
	}

	Log("Database connected")
	// return db

}

/* func ConnectD() *sql.DB {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "",
		Addr:                 "localhost:3306",
		DBName:               "go_database",
		AllowNativePasswords: true,
	}

	db, openErr := sql.Open("mysql", cfg.FormatDSN())

	if IsNotNil(openErr) {
		FLog(openErr)
	}

	if pingErr := db.Ping(); IsNotNil(pingErr) {
		FLog(pingErr)
	}

	Log("Database connected")
	return db
} */

/* func CloseDB(db *sql.DB) {
	defer db.Close()
	Log("database closed")
	os.Exit(1)
} */

func CloseDB() {
	defer db.Close()
	Log("database closed")
	os.Exit(1)
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Job      string `json:"job"`
}

// Pass username and job to create a new user
func Insert(db *sql.DB, user User) (int64, error) {
	defer Log(fmt.Sprintf("Insert new row for username: %v and job: %v",
		user.Username, user.Job))

	query := "INSERT INTO user(username, job) VALUES (?, ?)"

	result, err := db.Exec(query, user.Username, user.Job)

	if IsNotNil(err) {
		return 0, fmt.Errorf("error inserting 'user': %v", err)
	}

	id, err := result.LastInsertId()

	if IsNotNil(err) {
		return 0, fmt.Errorf("error getting last inserted Id: %v", err)
	}

	return id, nil
}

// Pass user's id to return the user record
func SelectById(db *sql.DB, id int64) (User, error) {
	defer Log(fmt.Sprintf("Select row by Id: %v", id))

	query := "SELECT `id`, `username`, `job` FROM `user` WHERE `id`=? LIMIT 1"

	row := db.QueryRow(query, id)

	// map the result onto the user object
	var _user User
	err := row.Scan(&_user.Id, &_user.Username, &_user.Job)

	if IsNotNil(err) && err == sql.ErrNoRows || IsNotNil(row.Err()) {
		return _user, fmt.Errorf(
			"error reading row by `id`: %v",
			err)
	}

	return _user, nil
}

// Return usr records
func SelectAll(db *sql.DB) ([]User, error) {
	defer Log("Select all rows")

	// map the result onto the user object
	var _users []User

	query := "SELECT `id`, `username`, `job` FROM `user`"

	rows, err := db.Query(query)

	if IsNotNil(err) || IsNotNil(rows.Err()) {
		return _users, fmt.Errorf("error reading rows of 'user': %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var _user User
		err = rows.Scan(&_user.Id, &_user.Username, &_user.Job)

		if IsNotNil(err) && err == sql.ErrNoRows {
			return _users, fmt.Errorf(
				"error reading row by `id`: %v",
				err)
		}

		_users = append(_users, _user)
	}

	return _users, nil
}

// Update user record, by reassigning the user property
// then call update to save the current state
// user.username="new username"
// user.job="new job"
// user.Update()
func UpdateById(db *sql.DB, user User) error {
	defer Log(fmt.Sprintf(
		"Update user by Id: %v, with username and jod as %v and %v",
		user.Id, user.Username, user.Job,
	))

	query := "UPDATE user SET username=?, job=? WHERE id=?"

	result, err := db.Exec(query, user.Username, user.Job, user.Id)

	if IsNotNil(err) {
		return fmt.Errorf("error updating row by Id: %v, %v", user.Id, err)
	}

	count, err := result.RowsAffected()

	if IsNotNil(err) {
		return fmt.Errorf("error getting last updated Id: %v", err)
	}

	if count < 1 {
		return fmt.Errorf("no row updated, rows affected: %v", count)
	}

	return nil
}

func DeleteById(db *sql.DB, id int64) error {
	defer Log(fmt.Sprintf("Delete user by Id: %v", id))

	query := "DELETE FROM user WHERE id=?"

	result, err := db.Exec(query, id)

	if IsNotNil(err) {
		return fmt.Errorf("error deleting row by Id: %v, %v", id, err)
	}

	count, err := result.RowsAffected()

	if IsNotNil(err) {
		return fmt.Errorf("error getting last updated Id: %v", err)
	}

	if count < 1 {
		return fmt.Errorf("no row updated, rows affected: %v", count)
	}

	return nil
}
