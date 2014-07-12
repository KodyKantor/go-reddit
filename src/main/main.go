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
	multiwriter := io.MultiWriter(file, os.Stdout)
	log := log.New(multiwriter, "Logger\t", log.Lshortfile)
	
	limit := 3
	sub := new(api.Subreddit)
	sub.Name = "gifs"
	links, err := sub.GetSub(log, api.NEW, limit)
	if err != nil {
		log.Println("Error getting sub", err)
		os.Exit(1)
	}
	
	for _, value := range links {
		fmt.Println(value)
	}
}
