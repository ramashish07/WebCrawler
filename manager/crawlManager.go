package manager

import (
	"fmt"
	"main/constant"
	"math"
	"time"

	"github.com/gocolly/colly/v2"
)

// Parallelism denotes the number of crawler workers
// RateLimit is the max speed per worker per hour
var (
	Parallelism=constant.INITIAL_PARALLELISM
	RateLimit=constant.INITIAL_RATE_LIMIT
)

// Function for crawling given URI
func CrawlPage(url string) (string, error) {
	

	time.Sleep(time.Second*10)

	//Checking if the result for URL is already stored in disk
    if !IsStoredInDisk(url) {
        html, err := CrawlPageRealTime(url)
        if err != nil {
            return "", err
        }

		fmt.Println("The link has not been found in the disk, crawling in real time ", url)
        
		StoreInDisk(url, html)
        
		return html, nil

    } else {
		fmt.Println("The link has been found in the disk, returning ", url)
	}

    cachedHTML, err := GetStoredPage(url)
    return cachedHTML, err
}

var startTime = time.Now()
var requestCounter = 0
// Result not stored in disk, crawl in real time using colly
func CrawlPageRealTime(url string) (string, error) {
	
	c := colly.NewCollector()
	
	
	// Limit the number of workers
	c.Limit(&colly.LimitRule{
        Parallelism: Parallelism,
    })

	var info string

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info=e.Text
		requestCounter++
	})
	
	err := c.Visit(url)
	if err != nil {
		fmt.Print(err)
	}
	
	// Calculating the crawling rate per hour and checking if it exceeds the limit
	elapsedTime := time.Since(startTime).Seconds()
	elapsedTime = math.Ceil(float64(elapsedTime/(3600)))
	requestRatePerHour := float64(requestCounter) / elapsedTime

	fmt.Println("The crawling speed is ", requestRatePerHour)

    if requestRatePerHour>float64(RateLimit){
		fmt.Println("The Rate Limit is Exceeded, The rate limit is ", RateLimit)
		fmt.Println("The Rate is ", requestRatePerHour)
		return "The Rate Limit is exceeded",err
	}
	
	return string(info),err
}
