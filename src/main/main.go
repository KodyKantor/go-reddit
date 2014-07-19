package main

import (
	"api"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	multiwriter := io.MultiWriter(file, os.Stdout)
	log := log.New(multiwriter, "Logger\t", log.Lshortfile)

	limit := 1
	sub := new(api.Subreddit)
	sub.Name = "gifs"

	//get the front page of a subreddit's TOP section
	page, err := sub.GetSub(log, api.TOP, 0, limit)
	if err != nil {
		log.Fatalln("Error getting sub:", err)
	}
	log.Println(page)

	comments, err := page.Top.GetComments(log, 1, 10)
	if err != nil {
		log.Fatalln("Error getting comments on link:", err)
	}
	
	for _, entry := range comments {
		log.Println(entry)
	}
	
	/*
		//get the next page
		page, err = sub.GetPage(log, api.NEXT)
		if err != nil {
			log.Fatalln("Error getting sub:", err)
		}
		log.Println(page)

		//get the previous page
		page, err = sub.GetPage(log, api.PREV)
		if err != nil {
			log.Fatalln("Error getting sub:", err)
		}
		log.Println(page)
	*/

}
