# WebCrawler
- Developed a high performance and robust web crawling service.
- Prioritized paying customers, ensuring faster access to web content,
and assigned more workers to them, along with implementing a retry
mechanism for improved performance.
- Implemented concurrent multi-threaded crawling with up to 5 workers
for paying customers and 2 for non-paying customers.
- Developed RESTful APIs for administrators to dynamically control the
crawl rate and to prevent crawl limit exceedance

## Table of Contents 
- [Tech Stack](#tech-stack)
- [Video Link](#video-link)
- [Installation](#installation)
- [Folder Structure](#folder-structure)
- [API Endpoints](#api-endpoints)
- [Models](#models)
- [Overview](#overview)
- [Features](#features)


## Tech Stack

*Client:* HTML, Bootstrap

*Server:* Golang

# Video Link 
https://drive.google.com/file/d/1fXgcFetgS1qRIfWAaVs8wvgTo9s8lNbE/view?usp=sharing
## Installation

### Prerequisites
1. Golang 
2. HTML 

To get started with the project, follow these steps:

1. Clone the repository:
   ```
   https://github.com/ramashish07/sendx-backend-iec2020095.git
   ```

2. Navigate to the project directory and run:
   ``` 
   go run main.go
   ```
   Note: Make sure all the library and modules are imported
  

## Folder Structure 
- The Folder Structure  follows MVC Architecture.
```
project-root/
├── constant/
    ├── constants.go
├── controller/
    ├── controller.go
    ├── customController.go
├── disk/
├── helper/
    ├── crawlHelper.go
├── manager/
    ├── crawlManager.go
    ├── diskManager.go
├── model/
    ├── models.go
├── router/
    ├── router.go
└── view/
    ├── index.html/
    
```

## API Endpoints 

#### Crawl an URL 

http
  POST:- /crawl  


| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `URL to be crawled` | `string` | *Required*. Your URL key
 |  `Request type`  |  `bool`    | User is paid or not 

#### Custom Crawl an URL 

http
  POST :- /customCrawl


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
	NumWorkers string 
	Rate string 
}`

3. Crawl Job
   
`type CrawlJob struct {
	URL string
	Priority int
}`

##  Overview 

-  When the user crawl  particular URL  without setting custom parameters  the  POST request goes into the server side.

- First the server checks the user is paid or non paid , if paid then it pushes the requested url into the paid  channel otherwise it pushes into the non paid channel.  Firstly the paid users url will be crawled giving priority to paid users.

- The server first  checks in the disk that the URL to be  crawled is stored in the disk or not in the last 60 min.

- If it is stored in the disk it returns the result from their otherwise it will do  real time crawling.

- If the real time crawling fails , the server will again retry for real time crawling upto three times.
  
- Once the crawling has been done the response is stored in the server disk , and after 60 minute time it will be deleted from the disk.

- For concurrent crawling , separate channels in golang had been made for paid and non paid users with workers assigned to them .

- For custom crawling , admin can set the number of workers and crawling rate per per worker per hour with the api  /customcrawl .  

- If the rate limit comes out to be more than the set value , the server will throw an error and inform that the rate limit  has been exceeded.




## Features 
- Priority given to the paid users in the first part because at first paid queue is crawled and also more workers assigned to them .
- Persistent Cache Storage that is disk is used here in the server side ,for that os module is used .
  ![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/96d14697-07ee-477b-87a4-cc11dba3e4bb)

- After a 60 min time after crawling had been done ,the response stored in the server disk will be deleted.
- Therefore if the URL crawled data is not in the disk , it will do real time crawling
![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/78b11f69-1051-48b4-8c07-cc287745a374)



- Retry Mechanism is implemented if the page is not available at that time.
- Crawled URL within a 60 min timeframe doesnot go for real time crawling instead their response are returned from disk.

  ![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/966be37b-3c96-4bc1-bcea-d5925cc01235)

- Workers assigned for paid workers and non paid workers which will concurrenly crawl URL for them ,more workers are assigned to paid workers.
- Custom Crawling has beed implemented in which the admin can set the  parallelism and rate limit per worker per hour.
![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/433b2959-974e-40ac-885b-a7afb186297d)

- The server throws notification if the crawl limit has been exceeded
 ![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/83f3982c-7c65-4de6-94cd-f45e9a377456)



- Error handled in each request and operation.
- Made a simple interface using html and bootstrap to make request from the user side and download crawled data in client local disk.

  ![image](https://github.com/ramashish07/sendx-backend-iec2020095/assets/91429764/dd80b762-66f5-4d22-aa9f-79953edb29f0)



  






  




