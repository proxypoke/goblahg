// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/goblahg
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Path      string // where the directory with the posts is located
	Intervall time.Duration // how often to reload posts, in minutes
}

func LoadConfig(path string) (conf Config) {
	log.Printf("Loading config from '%s'.\n", path)
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		log.Fatal(err)
	}
	conf.Intervall *= time.Minute

	log.Printf("Post directory is '%s'.\n", conf.Path)
	log.Printf("Reload intervall is '%q'.\n", conf.Intervall)
	return
}
