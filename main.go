// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/goblahg
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING
package main

import (
	"log"
	"net/http"
)

func main() {
	updates := make(chan Posts)
	// writes updates
	go WatchDir(".", updates)
	// reads updates
	go ServeBlog(updates)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
