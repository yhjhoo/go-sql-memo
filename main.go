package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Person struct {
	name string
	id   int64
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	query, err := db.Query("select name, id from person")
	for query.Next() {
		//var person Person
		//query.Scan(&person.name, &person.id)
		//println("ff: " + person.name + ", id: " + strconv.FormatInt(person.id, 20))

		var name string
		var id int64
		query.Scan(&name, &id)
		println(name + strconv.FormatInt(id, 20))
	}

}
