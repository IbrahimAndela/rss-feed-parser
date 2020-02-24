package scrape

import (
	"net/http"
	"encoding/xml"
	"io/ioutil"


	"github.com/mkn2016/rss-feed-parser/structs"
	"github.com/mkn2016/rss-feed-parser/errors"
	"github.com/mkn2016/rss-feed-parser/connection"
	"github.com/mkn2016/rss-feed-parser/config"
)

//URLScraper ... takes an argument of url string and returns an RSSParent Structure
func URLScraper(URL string) (structs.RssParent, error) {

	connected, connectionError := connection.IsConnected(config.GoogleURL)

	if connected {
		response, responseError := http.Get(URL)

		err := errors.CheckErr(responseError)

		defer response.Body.Close()

		if err {
			return structs.RssParent{}, responseError
		}
		parsedBytes, ioutilError := ioutil.ReadAll(response.Body)
		parsingError := errors.CheckErr(ioutilError)
		if parsingError {
			return structs.RssParent{}, ioutilError
		}
		rssParent := structs.RssParent{}
		xml.Unmarshal(parsedBytes, &rssParent)
		return rssParent, nil
	}
	return structs.RssParent{}, connectionError
}