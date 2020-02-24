package main

import (
	"log"
	// "time"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mkn2016/rss-feed-parser/schedule"
	"github.com/mkn2016/rss-feed-parser/db"
	"github.com/mkn2016/rss-feed-parser/apis"
)

func init() {
	db.CreateDatabase("RssFeeds")
	db.CreateTable("RssFeeds", "feeds")

	schedule.CronJobs()
}


func main(){
	log.Println("testing main")
	router := mux.NewRouter()
	router.HandleFunc("/feeds", apis.GetFeedsAPI).Methods("Get")
	router.HandleFunc("/feeds/{id:[a-zA-Z0-9]*}", apis.GetFeedByIDAPI).Methods("Get")
	log.Fatal(http.ListenAndServe(":8001", router))
}