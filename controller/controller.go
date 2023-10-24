package controller

import (
	"encoding/json"
	"fmt"
	"main/constant"
	"main/manager"
	"main/model"
	"net/http"
)

// Creating channels for paid, unpaid and completed requests
var (
	paidChannel   = make(chan string, constant.PAID_WORKERS) 
	unpaidChannel = make(chan string, constant.UNPAID_WORKERS) 
	crawledResult = make(chan string, constant.PAID_WORKERS+constant.UNPAID_WORKERS)
)

func StaticPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}

// Crawls pages in the specified channel
func crawlWorker(queue chan string) {
	for url := range queue {
		crawledHTML, err := manager.CrawlPage(url)
		if err != nil {
			crawledResult <- "Failed to crawl the page"
		} else {
			crawledResult <- crawledHTML
		}
	}
}

// API handler
func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The API is triggered")

	var request model.CrawlRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		fmt.Println("Failed to decode JSON request:", err)
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	url := request.URL
	isPayingCustomer := request.IsPayingCustomer

	fmt.Println("URL:", url)
	fmt.Println("Is Paying Customer:", isPayingCustomer)

	if isPayingCustomer {
		paidChannel <- url
	} else {
		unpaidChannel <- url
	}

	//Concurrently start crawling for paid and unpaid channels
	go crawlWorker(paidChannel)
	go crawlWorker(unpaidChannel)

    crawledHTML := <-crawledResult

    if crawledHTML == "Failed to crawl the page" {
        http.Error(w, "Failed to crawl the page", http.StatusInternalServerError)
    } else {
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        _ , _ = w.Write([]byte(crawledHTML))
    }
}