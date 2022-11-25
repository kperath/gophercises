package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLinkText(node *html.Node, linkText *string) {
	if node.Type == html.TextNode {
		*linkText += node.Data
	}
	for sib := node.FirstChild; sib != nil; sib = sib.NextSibling {
		getLinkText(sib, linkText)
	}
}

func addLink(node *html.Node) (link Link) {
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	getLinkText(node, &link.Text)
	return link
}

func traverse(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		*links = append(*links, addLink(node))
	}
	for sib := node.FirstChild; sib != nil; sib = sib.NextSibling {
		traverse(sib, links)
	}
}

type Link struct {
	Href string
	Text string
}

func getLinks(filename string) []Link {
	f, err := os.Open(filename)
	check(err)
	r := bufio.NewReader(f)
	node, err := html.Parse(r)
	check(err)
	var links []Link
	traverse(node, &links)
	return links
}

func main() {
	fmt.Println("==HTML PARSER==")
	if len(os.Args) != 2 {
		log.Fatal("html file argument required")
	}
	fileName := os.Args[1]
	fmt.Println(getLinks(fileName))
}
