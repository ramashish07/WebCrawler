package model

type CrawlRequest struct {
	URL              string
	IsPayingCustomer bool
}

type CrawlJob struct {
	URL      string
	Priority int
}