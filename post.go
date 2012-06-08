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
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Post struct {
	Content, Title string
	When           time.Time
}

// Naive basename implementation.
func basename(filename string) (basename string) {
	b := strings.Split(filename, "/")
	basename = b[len(b)-1]
	basename = strings.Split(basename, ".")[0]
	return
}

func FromFile(path string) (p Post, err error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	p.Content = string(blackfriday.MarkdownBasic(raw))
	p.Title = basename(path)
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	p.When = info.ModTime()
	return
}

func (p Post) Serve() (string, func(http.ResponseWriter, *http.Request)) {
	// TODO: Load from template
	handle := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "<html>")
		defer fmt.Fprintln(w, "</html>")

		fmt.Fprintln(w, "<head>")
		fmt.Fprintf(w, "<title>%s</title>", p.Title)
		fmt.Fprintln(w, "</header>")

		fmt.Fprintln(w, "<body>")
		defer fmt.Fprintln(w, "</body>")

		fmt.Fprint(w, p.Content)
	}
	return "/" + strings.ToLower(p.Title) + "/", handle
}
