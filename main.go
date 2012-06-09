// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/goblahg
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var confpath string
	flag.StringVar(&confpath, "config", "./blahgrc", "path to config")
	flag.Parse()

	conf := LoadConfig(confpath)

	updates := make(chan Posts)
	// writes updates
	go WatchDir(conf.Path, updates, conf.Intervall)

	posts := <-updates
	go http.HandleFunc("/", posts.Serve(conf.Path))

	// reads updates
	go ServeBlog(updates)

	updates <- posts

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
