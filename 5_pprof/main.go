package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var authTOken = flag.String("auth_token", "", "token to check Auth-Token header")

func main() {
	http.HandleFunc("/whats_time", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Auth-Token") == *authTOken {
			w.Write([]byte(time.Now().String()))
		}
		w.Write([]byte("304"))
	})

	http.ListenAndServe(":80", nil)
}

// curl --output - localhost:80/debug/pprof/cmdline

// use vault
