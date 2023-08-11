package main

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: Do validation

// Pass username and job to create a new user
func Insert(db *sql.DB, username, job string) (int64, error) {
	defer log.Printf("Insert new row for username: %v and job: %v\n", username, job)

	query := "INSERT INTO user(username, job) VALUES (?, ?)"

	result, err := db.Exec(query, username, job)

	if err != nil {
		return 0, fmt.Errorf("error inserting 'user': %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error getting last inserted Id: %v", err)
	}

	return id, nil
}

// Pass user's id to return the user record
func SelectById(db *sql.DB, id int64) (User, error) {
	defer log.Println("Select row by Id", id)

	query := "SELECT `id`, `username`, `job` FROM `user` WHERE `id`=? LIMIT 1"

	row := db.QueryRow(query, id)

	// map the result onto the user object
	var _user User
	err := row.Scan(&_user.Id, &_user.Username, &_user.Job)

	if err != nil && err == sql.ErrNoRows || row.Err() != nil {
		return _user, fmt.Errorf(
			"error reading row by `id`: %v",
			err)
	}

	return _user, nil
}

// Return usr records
func SelectAll(db *sql.DB) ([]User, error) {
	defer log.Println("Select all rows")

	// map the result onto the user object
	var _users []User

	query := "SELECT `id`, `username`, `job` FROM `user`"

	rows, err := db.Query(query)

	if err != nil || rows.Err() != nil {
		return _users, fmt.Errorf("error reading rows of 'user': %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var _user User
		err = rows.Scan(&_user.Id, &_user.Username, &_user.Job)

		if err != nil && err == sql.ErrNoRows {
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
func Update(db *sql.DB, user User) error {
	defer log.Printf(
		"Update user by Id: %v, with username and jod as %v and %v\n",
		user.Id, user.Username, user.Job,
	)

	query := "UPDATE user SET username=?, job=? WHERE id=?"

	_, err := db.Exec(query, user.Username, user.Job, user.Id)

	if err != nil {
		return fmt.Errorf("error updating row by Id: %v, %v", user.Id, err)
	}

	return nil
}

func Delete(db *sql.DB, id int64) error {
	defer log.Printf("Delete user by Id: %v\n", id)

	query := "DELETE FROM user WHERE id=?"

	_, err := db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error updating row by Id: %v, %v", id, err)
	}

	return nil
}
