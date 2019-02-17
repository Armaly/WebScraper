package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)
//Precondition: Is called by main
//Postconditions: Creates a database and relevant tables
func createDatabase(){
	//http://go-database-sql.org/overview.html
	//Creates database file
	fmt.Println("Creating database...")
	os.Create("overflow.db")

	db, err := sql.Open("sqlite3", "overflow.db")
	if err != nil{
		log.Fatal(err)
	}
	//Ping test
	err = db.Ping()
	if err != nil{
		fmt.Println("PING ERROR")
	}
	_, err =db.Exec("CREATE TABLE `Stack_Jobs` (`JOB_ID` INTEGER PRIMARY KEY AUTOINCREMENT, `TITLE` TEXT, `COMPANY` TEXT, `LOCATION` TEXT, `TAGS` TEXT, `DESCRIPTION` TEXT, `SALARY` TEXT NULL, `REMOTE` TEXT, `PUBLICATION_DATE` TEXT, `DATE_ADDED_TO_DB` TEXT)")
	if err!=nil{
		log.Fatal(err)
	}
	db.Close()
}
//Preconditions: Is called
//Postconditions: Populates the database with the scraped data from the custom struct array
func populateDatabase(list []*FeedStruct){

	db, err := sql.Open("sqlite3", "overflow.db")
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Populating database...")

	for i:= range list{

		statement := `INSERT INTO Stack_Jobs (JOB_ID, TITLE, COMPANY, LOCATION, TAGS, DESCRIPTION, REMOTE, PUBLICATION_DATE, DATE_ADDED_TO_DB)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

		_, err = db.Exec(statement, i, list[i].Title, list[i].Company, list[i].Location, list[i].Tags, list[i].Description, list[i].Remote, list[i].Date, list[i].DateAdded)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("Executed ",i)

	}
	fmt.Println("Database has been populated!")
}