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
