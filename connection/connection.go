package connection

import (
	"net"
	"errors"
)

// IsConnected ... checks whether we have an internet connection so as to pass the rss feeds from <url>
func IsConnected(url string) (bool, error) {
	url += ":http"
	_, err := net.Dial("tcp", url)
	if err != nil {
		return false, errors.New("Could not connect to the internet. \nRSS Feed Parser not initiated")
	}
	return true, nil
}