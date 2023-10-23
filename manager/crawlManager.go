package manager

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func CrawlPage(url string) (string, error) {
    if !IsStoredInDisk(url) {
        html, err := CrawlPageRealTime(url)
        if err != nil {
            return "", err
        }
		fmt.Println("The link has not been found in the cache ", url)
        
		StoreInDisk(url, html)
        
		return html, nil
    } else {
		fmt.Println("The link has been found in the cache ", url)
	}

    cachedHTML, err := GetStoredPage(url)
    return cachedHTML, err
}

func CrawlPageRealTime(url string) (string, error) {
	c := colly.NewCollector()
	var info string

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info=e.Text
	})
	
	err := c.Visit(url)
	if err != nil {
		fmt.Print(err)
	}

	return string(info),err
}