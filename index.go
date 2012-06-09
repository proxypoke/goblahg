// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/goblahg
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

type Posts []Post

func (posts *Posts) Len() int {
	return len(*posts)
}

// This is actually Greater, so it causes a reverse sort
func (posts *Posts) Less(i, j int) bool {
	return (*posts)[i].When.Unix() > (*posts)[j].When.Unix()
}

func (posts *Posts) Swap(i, j int) {
	(*posts)[i], (*posts)[j] = (*posts)[j], (*posts)[i]
}

// Return the n most recent posts, ordered by decreasing recency.
// (Most recent posts first)
func (p Posts) MostRecent(n int) (recent []Post) {
	sort.Sort(&p)
	for i := 0; i < n; n++ {
		recent = append(recent, p[i])
	}
	return
}

func (posts Posts) Contains(p Post) bool {
	for _, post := range posts {
		if p == post {
			return true
		}
	}
	return false
}

func LoadDir(path string) (posts Posts) {
	// the error doesn't matter, we'll just return no posts
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			post, err := FromFile(path + file.Name())
			if err != nil {
				log.Print(err)
				continue
			}
			posts = append(posts, post)
		}
	}
	return
}

func (posts Posts) Serve(path string) func(http.ResponseWriter, *http.Request) {
	handle := func(w http.ResponseWriter, req *http.Request) {
		posts = LoadDir(path)
		sort.Sort(&posts)

		fmt.Fprintln(w, "<html>")
		defer fmt.Fprintln(w, "</html>")

		fmt.Fprintln(w, "<head>")
		fmt.Fprintf(w, "<title>%s</title>", "Goblahg")
		fmt.Fprintln(w, "</head>")

		fmt.Fprintln(w, "<body>")
		defer fmt.Fprintln(w, "</body>")

		fmt.Fprint(w, "<h2>Most recent posts</h2>")

		fmt.Fprint(w, "<ul>")
		defer fmt.Fprint(w, "</ul>")

		for _, post := range posts {
			fmt.Fprint(w, "<li>")
			fmt.Fprintf(w, "<a href=/%s>%s</a>",
				strings.ToLower(post.Title), post.Title)
			fmt.Fprint(w, "</li>")
		}
	}
	return handle
}

func (posts Posts) Add(p Post) {
	posts = append(posts, p)
}
