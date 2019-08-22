package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sample/search"

	_"search"
)

type (
	item struct{
		XmlName xml.Name `xml:"item"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		Guid string `xml:"guid"`
	}
	image struct {
		XmlName xml.Name `xml:"image"`
		Url string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}

	channel struct {
		XmlName xml.Name `xml:"channel"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		PubDate string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		TTL string `xml:"ttl"`
		Generator string `xml:"generator"`
		Language string `xml:"language"`
		WebMaster string `xml:"webMaster"`
		Image image `xml:"image"`
		Item []item `xml:"item"`
	}

	rssDocument struct {
		XmlName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) retrieve(feed *search.Feed)(*rssDocument, error) {
	if feed.Url == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	resp, err := http.Get(feed.Url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document,err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string)([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s],Site[%s] For Url[%s]\n", feed.Type, feed.Site, feed.Url)

	document,err := m.retrieve(feed)
	if err != nil {
		return nil,err
	}
	for _,channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field : "Title",
				Content : channelItem.Title,
			})
		}

		matched,err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil,err
		}
		if matched {
			results = append(results, &search.Result{
				Field : "Description",
				Content : channelItem.Description,
			})
		}
	}
	return results,nil
}
