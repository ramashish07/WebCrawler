package router

import (
	"main/controller"

	"github.com/gorilla/mux"
)

// Declaring the routes for the app
func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controller.StaticPageHandler).Methods("GET")
	router.HandleFunc("/crawl", controller.CrawlHandler).Methods("POST")
	router.HandleFunc("/customCrawl",controller.CustomCrawlHandler).Methods("POST")
    return router
}