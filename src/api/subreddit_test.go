package api

import (
	"log"
	"testing"
)

func TestGetSub(t *testing.T) {
	//create a muted log
	log := log.New(new(TestWriter), "Testing", 0)

	s := new(Subreddit)
	//test for error return with bad subreddit
	s.Name = "asdfasdfasdf"
	if _, err := s.GetSub(log, HOT, 0, 100); err == nil {
		t.Fatal("Expecting error, received none")
	}

	//test for correct number of links returned
	s.Name = "gifs"
	page, err := s.GetSub(log, HOT, 0, 10)
	if err != nil {
		t.Fatal("Not expecting an error (sub should exist)")
	}
	if len(page.Links) != 10 {
		t.Fatal("Received incorrect number of links")
	}

	//test to see if page exists before the first link
	if _, err = s.GetSub(log, HOT, PREV, 10); err == nil {
		t.Fatal("Expecting error (no previous page should exist)")
	}

	//test for correct number of links on next page request
	page, err = s.GetSub(log, HOT, NEXT, 10)
	if err != nil {
		t.Fatal("Not expecting error (next page should exist)")
	}
	if len(page.Links) != 10 {
		t.Fatal("Received incorrect number of links")
	}

	//test bad http request string
	s.Name = "\n\n\n\n\n\n\n\n\n"
	if _, err = s.GetSub(log, HOT, 0, 10); err == nil {
		t.Fatal("Expecting error (bad http request string)")
	}

}

//a 'muted' logger to pass to the api
type TestWriter struct {
}

func (writer *TestWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
