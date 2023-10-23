package controller

import (
	"encoding/json"
	"fmt"
	"main/manager"
	"main/model"
	"net/http"
)


type PriorityQueue []model.CrawlJob
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(model.CrawlJob)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func StaticPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}


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


	priority := 0
    if isPayingCustomer  {
        priority= 1
    }

	piq := make(PriorityQueue, 0)
	piq.Push(model.CrawlJob{URL: url, Priority: priority})

	for piq.Len() > 0 {
		item := piq.Pop().(model.CrawlJob)
        crawledHTML, err := manager.CrawlPage(item.URL)
        if err != nil {
            http.Error(w, "Failed to crawl the page", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        _ , _ = w.Write([]byte(crawledHTML))
	}
}




