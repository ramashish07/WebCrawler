package controller

import (
	"encoding/json"
	"fmt"
	"main/manager"
	"main/model"
	"net/http"
	"strconv"
)

func CustomCrawlHandler(w http.ResponseWriter, r *http.Request) {
	var request model.CustomCrawlRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		fmt.Println("Failed to decode JSON request:", err)
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	numWorkers := request.NumWorkers
	crawlRate := request.Rate

	fmt.Println("Number of Workers:", numWorkers)
	fmt.Println("Crawl Rate:", crawlRate)

	numWorkersInt, err := strconv.Atoi(numWorkers)
	if err != nil {
		fmt.Println("Failed to convert numWorkers to an integer:", err)
		http.Error(w, "Invalid numWorkers value", http.StatusBadRequest)
		return
	}

	crawlRateInt, err := strconv.Atoi(crawlRate)
	if err != nil {
		fmt.Println("Failed to convert crawlRate to an integer:", err)
		http.Error(w, "Invalid crawlRate value", http.StatusBadRequest)
		return
	}

	manager.Parallelism=numWorkersInt
	manager.RateLimit=crawlRateInt

	w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
}
