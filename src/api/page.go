package api

import (
	"strconv"
)

//Page is a struct that represents the current page of a subreddit
type Page struct {
	Top     Link   //first link
	Bottom  Link   //last link
	Links   []Link //all links
	Section int    //Enum (HOT, NEW, TOP)
}

//String method for the Page type
//Prints all of the information about a page, including links
func (page Page) String() string {
	result := ""

	result += "Page contains " + strconv.Itoa(len(page.Links)) + " links\n"
	for _, value := range page.Links {
		result += value.String() + "\n"
	}

	return result
}