package api

import (
	"log"
	"strconv"
	"encoding/json"
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

func (link *Link) GetComments(log *log.Logger) {
	request := "http://www.reddit.com/r/" + link.Subreddit + "/comments/" + link.Id + ".json?depth=2&limit=2"
	log.Println("Request string is", request)
	body, err := ProcessRequest(request, "GET")

	if err != nil {
		log.Println("Error getting comments:", err)
	}
	
	type Listing []struct {
		Data struct {
			Children []struct {
				Kind string
				Data struct {
					Body string
					Edited bool
					Ups int
				}
			}
		}
	}
	log.Println("Json is", string(body))
	
	var listing Listing
	json.Unmarshal(body, &listing)
	log.Println("Json is", listing[1].Data.Children[0].Data.Body)
	
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
