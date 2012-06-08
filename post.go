package main

import (
	"fmt"
	//"github.com/russross/blackfriday"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Content, Title string
	when           time.Time
}

func (p Post) Serve() func(http.ResponseWriter, *http.Request) {
	// TODO: Load from template
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "<html>")
		defer fmt.Fprintln(w, "</html>")

		fmt.Fprintln(w, "<head>")
		fmt.Fprintf(w, "<title>%s</title>", p.Title)
		fmt.Fprintln(w, "</header>")

		fmt.Fprintln(w, "<body>")
		defer fmt.Fprintln(w, "</body>")

		fmt.Fprint(w, p.Content)
	}
}

/*
func main() {
	var p Post
	p.Title = "Foobar"
	p.Content = "Lorem ipsum etercerum"

	http.HandleFunc("/", p.Serve())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/
