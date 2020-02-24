package apis

import (
	// "github.com/mkn2016/rss-feed-parser/structs"
	// "log"
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"

	"github.com/mkn2016/rss-feed-parser/db"
)

//GetFeedsAPI ... Get all feeds from table and render to json
func GetFeedsAPI(w http.ResponseWriter, req *http.Request) {
	feeds := db.GetTableRecords("RssFeeds", "feeds")
	json.NewEncoder(w).Encode(feeds)
}

//GetFeedByIDAPI ... Get rss feed by id from table and render to json
func GetFeedByIDAPI(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	result := db.GetRecordByID("RssFeeds", "feeds", params["id"])
	json.NewEncoder(w).Encode(&result)
	w.Header().Add("Content-Type", "application/json")
}
