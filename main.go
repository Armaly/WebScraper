package main

import (
	"fmt"
)

func main(){

	fmt.Println("Scraping data...")
	list := PopulateStruct()
	fmt.Println("Done with scraping! Now creating/adding to database.")

	fmt.Println(list[2].Date)
	createDatabase()
	fmt.Println("Database and tables have been made!")

	populateDatabase(list)

	fmt.Println("Database populated!")
}