package schedule

import (
	"log"

	"gopkg.in/robfig/cron.v3"

	"github.com/mkn2016/rss-feed-parser/db"
	"github.com/mkn2016/rss-feed-parser/urlparser"
)

// CronJobs ... schedule cron job to be run after a specified interval in this case daily
func CronJobs() {
	c := cron.New()

	Task := func() {
		log.Println("Initiated Cron Job...")
		var urls []string

		urls = append(urls, "http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml?edition=int")
		urls = append(urls, "http://rss.cnn.com/rss/cnn_topstories.rss")

		doc := urlparser.ParseAllURLS(urls)
		
		db.InsertRecordToTable("RssFeeds", "feeds", doc)
	}
	
	c.AddFunc("*/1 * * * *", Task)
	c.Start()
	log.Println(c.Entries())
}