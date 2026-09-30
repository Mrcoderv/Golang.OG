// now calculating the page size by ussing the concurrency approach
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	size int
	Url  string
}

func responseUrlSize(url string, ch chan Page) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while getting the response from the url", err)
		ch <- Page{size: 0, Url: url}
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading the response body", err)
		ch <- Page{size: 0, Url: url}
		return
	}
	ch <- Page{size: len(body), Url: url}
}
func main() {
	pages := make(chan Page)

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.randommmmm.org",
		"https://www.github.com",
	}
	for _, url := range urls {
		go responseUrlSize(url, pages)
	}
	for range urls {
		page := <-pages
		fmt.Printf("URL: %s, Size: %d bytes\n", page.Url, page.size)
	}
}
