package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Id            int64
	Username, Job string
}

func main() {
	// connect to the database using rhe following configuration
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "",
		Addr:                 "localhost:3306",
		DBName:               "go_database",
		AllowNativePasswords: true,
	}

	// get a database handle
	db, openErr := sql.Open("mysql", cfg.FormatDSN())

	if openErr != nil {
		log.Fatal(openErr)
	}

	// ping database to get connection status
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	defer func() {
		log.Println("database closed")
		db.Close()
	}()

	log.Println("Database connected")

	// create user
	/* id, insertError := Insert(db, "john doe", "Software engineer")

	if insertError != nil {
		log.Println(insertError)
	} else {
		log.Println("User created with ID:", id)
	} */

	// read user by Id
	/* user, err := SelectById(db, 1)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("User {Id: %v, Username: %v, Job: %v}\n", user.Id, user.Username, user.Job)
	} */

	// read all users
	/* if users, selectAllErr := SelectAll(db); selectAllErr != nil {
		log.Println(selectAllErr)
	} else {
		for _, user := range users {
			log.Printf("User {Id: %v, Username: %v, Job: %v}\n", user.Id, user.Username, user.Job)
		}
	} */

	// update user by id
	/* updateErr := Update(db, User{
		Id:       1,
		Username: "Daniel",
		Job:      "Cyber Security Expert",
	})

	if updateErr != nil {
		log.Println(updateErr)
	} */

	// delete
	/* if deleteErr := Delete(db, 1); deleteErr != nil {
		log.Println(deleteErr)
	} */

}
