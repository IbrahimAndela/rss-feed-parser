package main

import (
	"log"
	// "time"
	// "http"

	// "github.com/mkn2016/rss-feed-parser/schedule"
	"github.com/mkn2016/rss-feed-parser/db"
	// "github.com/mkn2016/rss-feed-parser/urlparser"
)

func init() {
	// db.CreateDatabase("RssFeeds")
	// db.CreateTable("RssFeeds", "feeds")

	// schedule.CronJobs()
	// log.Println("Initiated Cron Job...")
	// var urls []string

	// urls = append(urls, "http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml?edition=int")
	// urls = append(urls, "http://rss.cnn.com/rss/cnn_topstories.rss")

	// doc := urlparser.ParseAllURLS(urls)
	// log.Println(doc)
	
	// db.InsertRecordToTable("RssFeeds", "feeds", doc)

	// time.Sleep(10*time.Second)
	// ret := db.GetTableRecords("RssFeeds", "feeds")
	// log.Println(ret)
	result := db.GetRecordByID("RssFeeds", "feeds", "edcc3489-54ed-466f-bad3-aadc26e5be11")
	log.Println(result)
}

func main(){
	log.Println("testing main")

	// schedule.CronJobs()
	// log.Println("Initiated Cron Job...")
	// var urls []string

	// urls = append(urls, "http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml?edition=int")
	// urls = append(urls, "http://rss.cnn.com/rss/cnn_topstories.rss")

	// doc := urlparser.ParseAllURLS(urls)
	// log.Println(doc)
	
	// db.InsertRecordToTable("RssFeeds", "feeds", doc)

	// time.Sleep(10*time.Second)
	// ret := db.GetTableRecords("RssFeeds", "feeds")
	// log.Println(ret)
	// log.Printf("%T", ret)
}