package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// www.s.com/a/b/c/a
// {s.com/, s.com/a, s.com/a/b, s.com/a/b/c}

type Set map[string]struct{}

func (s Set) Add(k string) {
	s[k] = struct{}{}
}

func getLinks(node *html.Node, currentPath string, baseURL *url.URL, visited Set, depth int) {
	if depth == 0 {
		return
	}

	if node.Type == html.ElementNode && node.Data == "a" {
		// found link
		for _, a := range node.Attr {
			if a.Key == "href" {
				link, err := url.Parse(a.Val)
				if err != nil {
					break
				}
				if link.Host != "" && link.Host != baseURL.Host {
					break
				}

				if _, ok := visited[currentPath+link.Path]; ok {
					break
				}

				nextLink := currentPath + link.Path
				resp, err := http.Get(nextLink)
				check(err)
				defer resp.Body.Close()
				data, err := io.ReadAll(resp.Body)
				check(err)
				doc, err := html.Parse(strings.NewReader(string(data)))
				check(err)
				visited.Add(nextLink)
				fmt.Println(visited)
				getLinks(doc, nextLink, baseURL, visited, depth-1)
				break
			}
		}
	}
	for sib := node.FirstChild; sib != nil; sib = sib.NextSibling {
		getLinks(sib, currentPath, baseURL, visited, depth-1)
	}
}

func main() {
	var domain string = "https://kperath.com"
	resp, err := http.Get(domain)
	check(err)
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	check(err)
	u, err := url.Parse(domain)
	check(err)

	s := make(Set)
	getLinks(doc.FirstChild, u.Host, u, s, 2)
	fmt.Println(s)
}
