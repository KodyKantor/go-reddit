package api

import (
	"encoding/json"
	"log"
	"strconv"
)

//Link is a struct that holds information about a specific
//link on a page within a subreddit
type Link struct {
	Title     string
	Score     int
	Domain    string //either self.Title or a web domain
	Url       string //url to the link
	Name      string //Fullname
	Author    string
	Ups       int
	Downs     int
	Created   float64 //Date created
	Id        string
	Subreddit string
}

// GetComments returns a slice of top-level comments on a link
// TODO implement a way to retrieve sub-comments
// this means that 'depth' is only supported for values <= 1
func (link *Link) GetComments(log *log.Logger, depth, limit int) (comments []Comment, err error) {
	if depth < 1 {
		depth = 10
	}
	if limit < 1 {
		limit = 10
	}
	request := "http://www.reddit.com/r/" + link.Subreddit + "/comments/" + link.Id + ".json?depth=" + strconv.Itoa(depth) + "&limit=" + strconv.Itoa(limit)
	log.Println("Request string is", request)
	body, err := ProcessRequest(request, "GET")

	if err != nil {
		return nil, err
	}

	type Listing []struct {
		Data struct {
			Children []struct {
				Kind string
				Data Comment
			}
		}
	}
	log.Println("Json is", string(body))

	var listing Listing
	json.Unmarshal(body, &listing)

	comments = make([]Comment, len(listing[1].Data.Children))
	for i, entry := range listing[1].Data.Children {
		comments[i] = entry.Data
	}

	return comments, err
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
	result += "Fullname is: \t" + link.Name + "\n"
	result += "Id is: \t" + link.Id

	return result
}
