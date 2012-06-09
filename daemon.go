// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/goblahg
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Regularly load a directory and write the posts in the given channel.
func WatchDir(path string, writeback chan Posts) {
	// TODO: reloading updates of existing posts
	var known Posts
	for {
		log.Printf("Reloading directory '%s'.\n", path)
		posts := LoadDir(path)
		var updates Posts
		for _, post := range posts {
			if !known.Contains(post) {
				updates = append(updates, post)
				known = append(known, post)

			}
		}
		writeback <- updates
		time.Sleep(30 * time.Second)
	}
}

func LoadDir(path string) (posts Posts) {
	// the error doesn't matter, we'll just return no posts
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			post, err := FromFile(file.Name())
			if err != nil {
				log.Print(err)
				continue
			}
			posts = append(posts, post)
		}
	}
	return
}

func ServeBlog(updates chan Posts) {
	for {
		posts := <-updates
		for _, post := range posts {
			log.Printf("Registering post '%s'.\n", post.Title)
			http.HandleFunc(post.Serve())
		}
	}
}
