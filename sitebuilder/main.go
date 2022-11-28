package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var domain string = "https://kperath.com"

func normalizeDomain(domain string) string {
	prefixes := []string{"https://", "http://", "www."}
	for _, p := range prefixes {
		domain = strings.TrimPrefix(domain, p)
	}
	return domain
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		// found link
		for _, a := range node.Attr {
			if a.Key == "href" {
			}
		}
	}
	for sib := node.FirstChild; sib != nil; sib = sib.NextSibling {
		getLinks(node)
	}
}

func main() {
	resp, err := http.Get(domain)
	check(err)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	check(err)
	doc, err := html.Parse(strings.NewReader(string(data)))
	check(err)
	getLinks(doc)
}
