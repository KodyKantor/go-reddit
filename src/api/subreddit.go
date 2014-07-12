package api

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
)

type Subreddit struct {
	Name string //name of subreddit
}

const (
	HOT int = iota
	NEW int = iota
	TOP int = iota
)

type Link struct {
	Title string
	Score int
	Domain string
	Url string
}

func (s *Subreddit) GetName(log log.Logger) string {
	return s.Name
}

func (s *Subreddit) GetSub(log *log.Logger, section int, limit int) (links []Link, err error) {
	if limit < 0 {
		limit = 10
	}
	
	var sec string
	switch section {
		case HOT: sec = "hot"
		case NEW: sec = "new"
		case TOP: sec = "top"
	}
	str := "http://www.reddit.com/r/" + s.Name + "/" + sec + ".json?limit=" + strconv.Itoa(limit)
	log.Println("String is", str)
	
	resp, err := http.Get(str)
	if err != nil {
		log.Println("Error connecting", err)
		return nil, err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading", err)
		return nil, err
	}
	
	type Listing struct {
		Data struct {
			Children []struct {
				Data Link 
			}
		}
	}
	
	var listing Listing
	json.Unmarshal(body, &listing)
	
	responses := listing.Data.Children
	count := len(responses)
	log.Printf("Received %d responses.", count)
	
	retval := make([]Link, count, count)
	
	for i, entry := range responses {
		retval[i] = entry.Data
	}
	
	return retval, nil
}

func (link Link) String() string {
	result := ""
	
	result += "Title is: \t" + link.Title + "\n"
	result += "Score is: \t" + strconv.Itoa(link.Score) + "\n"
	result += "URL is: \t" + link.Url + "\n"
	result += "Domain is: \t" + link.Domain
	
	return result
}