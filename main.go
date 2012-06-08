package main

import (
	"log"
	"net/http"
)

func main() {
	p, err := FromFile("foobar.md")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc(p.Serve())
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
