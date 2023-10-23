package router

import (
	"main/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controller.StaticPageHandler).Methods("GET")
	router.HandleFunc("/crawl", controller.CrawlHandler).Methods("POST")
    return router

}