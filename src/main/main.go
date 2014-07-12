package main

import (
	"fmt"
	"api"
	"log"
	"os"
	"io"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error creating file", err)
		os.Exit(1)
	}
	defer file.Close()
	
	multiwriter := io.MultiWriter(file, os.Stdout)
	log := log.New(multiwriter, "Logger\t", log.Lshortfile)
	
	limit := 10
	sub := new(api.Subreddit)
	sub.Name = "gifs"
	page, err := sub.GetSub(log, api.TOP, "", limit)
	if err != nil {
		log.Fatalln("Error getting sub", err)
	}
	
	log.Println(page)

	
	page, err = sub.NextPage(log)
	
	if err != nil {
		log.Fatalln("Error getting sub", err)
	}
	
	log.Println(page)

}
