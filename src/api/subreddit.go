/*
The GoLang Reddit API wrapper
The api package wraps the REST Reddit API
with user-friendly Go function calls.
*/
package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Enum that represents the sections within a subreddit, or
//next-page, prev-page
const (
	HOT int = iota
	NEW int = iota
	TOP int = iota

	NEXT int = iota
	PREV int = iota

	AGENT = "GoLang Reddit API by /u/kantosaurus"
)

//Subreddit is a struct that represents a page of links
//related to a specific topic
type Subreddit struct {
	Name string //name of subreddit
	Page Page   //the page that we are currently viewing
}

//GetSub will query Reddit's servers for the JSON of a target subreddit.
func (s *Subreddit) GetSub(log *log.Logger, section int, place int, limit int) (page Page, err error) {
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

	switch place {
	case NEXT:
		str += "&after=" + s.Page.Bottom.Name
		log.Println("Retrieving the next page")
	case PREV:
		str += "&before=" + s.Page.Top.Name
		log.Println("Retrieving the previous page")
	default:
	}
	log.Println("Request string is", str)

	client := http.Client{}
	req, err := http.NewRequest("GET", str, nil)
	req.Header.Set("User-Agent", AGENT)
	resp, err := client.Do(req)

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
	if count < 1 {
		err = errors.New("No links received")
		return Page{}, err
	}
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
func (s *Subreddit) GetPage(log *log.Logger, place int) (page Page, err error) {
	log.Println("Getting next page")
	return s.GetSub(log, s.Page.Section, place, len(s.Page.Links))
}
