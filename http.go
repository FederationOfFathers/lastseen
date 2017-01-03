package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/FederationOfFathers/consul"
)

var listenOn = ":8890"

func init() {
	flag.StringVar(&listenOn, "listen", listenOn, "ip:port to listen on for requests")
}

func mindHTTP() {
	consul.Must()
	if err := consul.RegisterOn("last-seen", listenOn); err != nil {
		panic(err)
	}
	http.HandleFunc("/seen.json", func(w http.ResponseWriter, r *http.Request) {
		e := json.NewEncoder(w)
		e.Encode(seen)
	})
	go http.ListenAndServe(listenOn, nil)
}
