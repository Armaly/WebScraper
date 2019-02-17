package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"regexp"
	"strings"
	"time"
)

//Preconditions: Is called by main
//Postconditions: Outputs a populated list struct array with the strings parsed from the url
func PopulateStruct()[]*FeedStruct{

	//Creates a new feed and sets it to the stackoverflow url with the set location and distance parameters
	//https://github.com/mmcdole/gofeed#basic-usage
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://stackoverflow.com/jobs/feed")

	//Creates Feedstruct array
	var list []*FeedStruct

	currentTime := time.Now()

	for i := range feed.Items{
		temp := &FeedStruct{
			Title: getTitle(feed.Items[i].Title),
			Company: getCompany(feed.Items[i].Title),
			Location: getLocation(feed.Items[i].Title),
			Tags: getCategories(feed.Items[i].Categories),
			//Salary
			Remote: getRemote(feed.Items[i].Title, feed.Items[i].Description),
			Date: feed.Items[i].Published,
			Url:  feed.Items[i].Link,
			//http://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
			DateAdded: currentTime.Format("06-Jan-02"),
			Description: feed.Items[i].Description,
			Job_Id: i,
		}

		//Places parsed strings into array
		list = append(list, temp)
	}
	return list
}
//Preconditions: Is called
//Postconditions: Parses string and returns the location pattern it separates everything between the (, and )
func getLocation(title string)string{

	r, _ := regexp.Compile(`(.*)\((.*[A-z]*), ([A-z]*)\)`)
	match := r.FindAllStringSubmatch(title, -1)

	if match == nil{
		fmt.Println("Error LOCATION found and handled!") //This is for that 1 COMPANY in Vienna that won't follow the regular naming conventions
		r, _ := regexp.Compile(`\(([^\)]+)\)`)
		match := r.FindAllStringSubmatch(title, -1)
		return match[0][1]
	}
	return match[0][2]
}
//Preconditions: Is called
//Postconditions: Parses string and returns the job title pattern it separates everything between the (, and )
func getTitle(title string)string{
	r, _ := regexp.Compile(`(.*)\((.*[A-z]*), ([A-z]*)\)`)
	match := r.FindAllStringSubmatch(title, -1)

	if match == nil{
		fmt.Println("Error TITLE found and handled!") //This is for that 1 COMPANY in Vienna that won't follow the regular naming conventions
		return title
	}
	return match[0][1]
}
//Preconditions: Is Called
//Postconditions: Parses the string and returns the company pattern is looks for "at" in string then takes everything up to the ( which is the beginning of the location
func getCompany(title string)string{
	r, _ := regexp.Compile(`at(.*)\(`)
	match := r.FindAllStringSubmatch(title, -1)

	return match[0][1]
}
//Preconditions: Is called by keyWordMatch
//Postconditions: Returns boolean depending on if the keyWord is in the string
func parseDescription(keyWord string, text string)bool{

	fmt.Println("STRING ACQUIRED")

	return strings.Contains(strings.ToUpper(keyWord), strings.ToUpper(text))
}
//Preconditions: Is called
//Postconditions: Returns a yes or no depending on if remote is available(SQLite does not support booleans)
func getRemote(title string, description string)string{
	if(strings.Contains(title, "remote") || strings.Contains(description, "remote")){
		return "YES"
	}
	return "NO"
}
//Preconditions: Is called
//Postconditions: Combines the categories string array into a single array and returns
func getCategories(categories []string)string{
	var combinedArray string

	for i:= range categories{
		combinedArray += categories[i] +", "
	}

	return combinedArray
}
