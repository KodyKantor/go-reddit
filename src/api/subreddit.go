/*
The GoLang Reddit API wrapper
The api package wraps the REST Reddit API
with user-friendly Go function calls.
*/
package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Enum that represents the sections within a subreddit
const (
	HOT int = iota
	NEW int = iota
	TOP int = iota
)

//Subreddit is a struct that represents a page of links
//related to a specific topic
type Subreddit struct {
	Name string //name of subreddit
	Page Page   //the page that we are currently viewing
}

//Page is a struct that represents the current page of a subreddit
type Page struct {
	Top     Link   //first link
	Bottom  Link   //last link
	Links   []Link //all links
	Section int    //Enum (HOT, NEW, TOP)
}

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

//GetSub will query Reddit's servers for the JSON of a target subreddit.
func (s *Subreddit) GetSub(log *log.Logger, section int, after string, limit int) (page Page, err error) {
	if limit < 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	var sec string
	switch section {
	case HOT:
		sec = "hot"
	case NEW:
		sec = "new"
	case TOP:
		sec = "top"
	default:
		sec = "hot"
	}

	log.Printf("User is requesting %d '%s' articles.", limit, sec)
	str := "http://www.reddit.com/r/"
	str += s.Name + "/"
	str += sec + ".json"
	str += "?limit=" + strconv.Itoa(limit)
	if after != "" {
		str += "&after=" + after
	}
	log.Println("Request string is", str)

	resp, err := http.Get(str)
	if err != nil {
		log.Println("Error connecting:", err)
		return Page{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return Page{}, err
	}

	//A struct for parsing the repsonse JSON
	type Listing struct {
		Data struct {
			Children []struct {
				Data Link
			}
		}
	}
	log.Println("Json is", string(body))
	var listing Listing
	json.Unmarshal(body, &listing)

	responses := listing.Data.Children
	count := len(responses)
	log.Printf("Received %d links.", count)

	links := make([]Link, count)

	//Place the links in a slice
	for i, entry := range responses {
		links[i] = entry.Data
	}

	s.Page = Page{links[0], links[len(links)-1], links, section}

	return s.Page, nil
}

//NextPage will use the GetSub() function to get the next page
//of a subreddit
func (s *Subreddit) NextPage(log *log.Logger) (page Page, err error) {
	log.Println("Getting next page")
	return s.GetSub(log, s.Page.Section, s.Page.Bottom.Name, len(s.Page.Links))
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
