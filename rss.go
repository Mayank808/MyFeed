package main

import (
	"encoding/xml"
	// "fmt"
	"io"
	"net/http"
	"time"
)

const (
	requestTimeout = 10 * time.Second
)

func rssURLToSocialFeed(url string) (RSSSocialFeed, error) {
	httpClient := http.Client{
		Timeout: requestTimeout,
	}

	response, err := httpClient.Get(url)
	if err != nil {
		return RSSSocialFeed{}, err
	}

	defer response.Body.Close()

	rssData, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSSocialFeed{}, err
	}

	rssSocialFeed := RSSSocialFeed{}

	if err = xml.Unmarshal(rssData, &rssSocialFeed); err != nil {
		return RSSSocialFeed{}, err
	}

	// fmt.Printf("%#v\n", rssSocialFeed)

	return rssSocialFeed, nil
}

// Models for parsing XML RSS feeds
type RSSSocialFeed struct {
	Channel struct {
		Title          string      `xml:"title" json:"title"`
		Description    string      `xml:"description" json:"description"`
		Link           string      `xml:"link" json:"link"`
		Language       string      `xml:"language" json:"language"`
		Item           []RSSItem   `xml:"item" json:"items"`
		Image          Image       `xml:"image" json:"image"`
		ItunesOwner    ItunesOwner `xml:"owner" json:"itunes_owner"`
		ItunesAuthor   string      `xml:"author" json:"itunes_author"`
		ItunesCategory string      `xml:"category" json:"itunes_category"`
	} `xml:"channel"`
}

type RSSItem struct {
	// Common fields
	Title           string       `xml:"title" json:"title"`
	Description     string       `xml:"description" json:"description"`
	Link            string       `xml:"link" json:"link"`
	PublicationDate string       `xml:"pubDate" json:"publication_date"`
	MediaContent    MediaContent `xml:"media:content" json:"media_content"` // News feeds
	ItunesDuration  string       `xml:"duration" json:"itunes_duration"`
	ItunesAuthor    string       `xml:"author" json:"itunes_author"`
	ItunesExplicit  string       `xml:"explicit" json:"itunes_explicit"`
	ItunesSummary   string       `xml:"summary" json:"itunes_summary"`
	ItunesSubtitle  string       `xml:"subtitle" json:"itunes_subtitle"`
	ItunesImage     Image        `xml:"image" json:"itunes_image"`
}

type MediaContent struct {
	Height string `xml:"height,attr" json:"height"`
	URL    string `xml:"url,attr" json:"url"`
	Width  string `xml:"width,attr" json:"width"`
}

type Image struct {
	Title string `xml:"title" json:"title"`
	URL   string `xml:"url" json:"url"`
	Link  string `xml:"link" json:"link"`
	Href  string `xml:"href,attr" json:"href"`
}

type ItunesOwner struct {
	Name  string `xml:"itunes:name" json:"name"`
	Email string `xml:"itunes:email" json:"email"`
}
