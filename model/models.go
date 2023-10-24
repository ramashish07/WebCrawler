package model

// Models for API-requests

type CrawlRequest struct {
	URL              string
	IsPayingCustomer bool
}

type CustomCrawlRequest struct {
	NumWorkers string `json:"numWorkers"`
	Rate       string `json:"crawlRate"`
}

type CrawlJob struct {
	URL      string
	Priority int
}

// Custom Queue Implementation

type Queue struct {
	item_value []string
}

func (q *Queue) Enqueue(item string) {
	q.item_value = append(q.item_value, item)
}

func (q *Queue) Dequeue() string {
	item := q.item_value[0]
	q.item_value = q.item_value[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.item_value) == 0
}