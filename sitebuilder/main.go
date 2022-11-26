package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/kperath/gophercises/sitebuilder/link"
	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	resp, err := http.Get("https://kperath.com")
	check(err)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(data))
	node, err := html.Parse(strings.NewReader(string(data)))
	check(err)
}
