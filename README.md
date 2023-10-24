# SendX Backend Assignment 

I have implemented all the three parts Required,Good to have and Great to have.



## Tech Stack

*Client:* HTML, Bootstrap

*Server:* Golang,Colly  




## API Endpoints 

#### Crawl an URL 

http
  POST /crawl  


| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `URL to be crawled` | `string` | *Required*. Your API key
 |  `Request type`  |  `bool`    | User is paid or not 

#### Custom Crawl an URL 

http
  POST  /customCrawl


| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|  Rate Limit| `int`  |  Rate limit per hour per worker |
|  Workers   | `int`    | Number of worker to be crawled concurrently 




## Models 
1. CrawlRequest
  
 `type CrawlRequest struct {
	URL string
	IsPayingCustomer bool
}`

2. Customer Crawl request
    
`type CustomCrawlRequest struct {
	NumWorkers string `json:"numWorkers"
	Rate       string `json:"crawlRate"
}`

3. Crawl Job
   
`type CrawlJob struct {
	URL      string
	Priority int
}`

##  Overview 

-  When the user crawl  particular URL  without setting custom paremeters  the  POST request goes into the server side.

- First the server checks the user is paid or non paid , if paid then it pushes the requested url into the paid  channel otherwise it pushes into the non paid channel.  Firstly the paid userls url will be crawled giving priority to paid users.

- The server first  checks in the disk that the URL to be  crawled is stored in the disk or not in the last 60 min.

- If it is stored in the disk it returns the result from their otherwise it will do  real time crawling.

- If the real time crawling fails , the server will again retry for real time crawling upto three times.

- For concurrent crawling , separate channels in golang had been made for paid and non paid users with workers assigned to them .

- For custom crawling , adming can set the number of workers and crawling rate per per worker per hour with the api  /customcrawl .  

- If the rate limit comes out to be more than the set value , the server will throw an error and inform that the rate limit  has been exceeded.



## Folder Structure 
- The Folder Structure  follwes MVC Architecture.
![image](https://github.com/ramashish07/chatgpt/assets/91429764/f06648f0-d6f8-4a4c-a013-61ad6f504199)



  




