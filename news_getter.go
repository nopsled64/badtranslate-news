package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// RSS is this
type RSS struct {
	XMLName  xml.Name  `xml:"rss"`
	Channels []Channel `xml:"channel"`
}

//Channel is this
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

// Item : this
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
}

//function inspired by https://tutorialedge.net/golang/parsing-xml-with-golang/
func getNewsRSS() (RSS, error) {
	var rss RSS
	// alt feeds.bbci.co.uk/news/technology/rss.xml
	if data, err := getContent("http://feeds.bbci.co.uk/news/politics/rss.xml"); err != nil {
		log.Printf("Failed to get XML: %v", err)
		return rss, err
	} else {
		log.Println("Received XML:")
		log.Println(string(data))
		xml.Unmarshal(data, &rss)
	}

	return rss, nil
}

//function taken from https://stackoverflow.com/questions/42717716/reading-content-from-http-get-in-golang
func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
