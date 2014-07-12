package api

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
)

const ( //enum for choosing subreddit section
	HOT int = iota
	NEW int = iota
	TOP int = iota
	RANDOM int = iota
)

type Subreddit struct {
	Name string //name of subreddit
	Page Page //the page that we are currently viewing
}

type Page struct {
	Top Link //first link
	Bottom Link //last link
	Links []Link //all links
	Section int
}

type Link struct {
	Title string
	Score int
	Domain string
	Url string
	Name string
	Author string
	Ups int
	Downs int
	Created float64
}

func (s *Subreddit) GetSub(log *log.Logger, section int, after string, limit int) (page Page, err error) {
	if limit < 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	
	var sec string
	switch section {
		case HOT: sec = "hot"
		case NEW: sec = "new"
		case TOP: sec = "top"
		default: sec = "hot"
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
	log.Printf("Received %d responses.", count)
	
	links := make([]Link, count, count)
	
	for i, entry := range responses {
		links[i] = entry.Data
	}
	
	s.Page = Page{links[0], links[len(links)-1], links, section}
	
	
	return s.Page, nil
}

func (s *Subreddit) NextPage(log *log.Logger) (page Page, err error) {
	log.Println("Getting next page")
	return s.GetSub(log, s.Page.Section, s.Page.Bottom.Name, len(s.Page.Links))
}

func (link Link) String() string {
	result := ""
	
	result += "Title is: \t" + link.Title + "\n"
	result += "Score is: \t" + strconv.Itoa(link.Score) + "\n"
	result += "Upvotes: \t" + strconv.Itoa(link.Ups) +"\n"
	result += "Downvotes: \t" + strconv.Itoa(link.Downs) + "\n"
	result += "URL is: \t" + link.Url + "\n"
	result += "Domain is: \t" + link.Domain + "\n"
	result += "Fullname is: \t" + link.Name
	
	return result
}

func (page Page) String() string {
	result := ""
	
	result += "Page contains " + strconv.Itoa(len(page.Links)) +" links\n"
	for _, value := range page.Links {
		result += value.String() + "\n"
	}
	
	return result
}