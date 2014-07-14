package api

import (
	"strconv"

)

//Link is a struct that holds information about a specific
//link on a page within a subreddit
type Link struct {
	Title   string
	Score   int
	Domain  string //either self.Title or a web domain
	Url     string //url to the link
	Name    string //Fullname
	Author  string
	Ups     int
	Downs   int
	Created float64 //Date created
}

//String method for the Link type
//Prints all of the information about a link
func (link Link) String() string {
	result := ""

	result += "Title is: \t" + link.Title + "\n"
	result += "Score is: \t" + strconv.Itoa(link.Score) + "\n"
	result += "Upvotes: \t" + strconv.Itoa(link.Ups) + "\n"
	result += "Downvotes: \t" + strconv.Itoa(link.Downs) + "\n"
	result += "URL is: \t" + link.Url + "\n"
	result += "Domain is: \t" + link.Domain + "\n"
	result += "Fullname is: \t" + link.Name

	return result
}

