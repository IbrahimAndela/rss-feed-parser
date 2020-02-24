package urlparser

import (
	"log"
	
	"github.com/mkn2016/rss-feed-parser/structs"
	"github.com/mkn2016/rss-feed-parser/scrape"
)

//ParseAllURLS ... parses all the urls in urls and appends to array of type RssParent
func ParseAllURLS(urls []string) []structs.RssParent{
	jobs := make(chan structs.RssParent, len(urls))
	done := make(chan bool)
	errs := make(chan error, len(urls))
	
	for _, url := range urls {
		go func(urlstr string) {
			b, err := scrape.URLScraper(urlstr)
			if err != nil {
				errs <- err
				jobs <- structs.RssParent{}
				done <- false
			}
			errs <- nil
			log.Println("Appending job from : ", urlstr)
			jobs <- b
			log.Println("Successfully added job from: ", urlstr, "to jobs channel")
			done <- true
			log.Println("Completed all jobs")
		}(url)
	}
	taskQueue := make([]structs.RssParent, 0)
	for i := 0; i < len(urls); i++ {
		taskQueue = append(taskQueue, <-jobs)
	}
	return taskQueue
}