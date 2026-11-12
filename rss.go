package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// there are two common feed formats
// rss 2.0 and atom
// this version works for rss 2.0

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func URLtoFeed(url string) (RSSFeed,error) {
	// create a http client 
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// get a response from that client 
	res,err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{},err
	}
	defer res.Body.Close()

	// res will get the data in binary from the url so dat will read its body
	dat,err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{} , err 
	}

	// we need to read all the binaries into the created structs 
	rssfeed := RSSFeed{}

	err = xml.Unmarshal(dat,&rssfeed)
	if err != nil {
		return RSSFeed{},err
	}

	// return the struct 
	return rssfeed , nil 
}