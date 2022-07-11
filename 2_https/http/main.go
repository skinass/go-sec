package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/whats_time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	http.ListenAndServe(":80", nil)
}

// sudo tcpdump -A -i lo0 'port 80'
// curl http://localhost:80/whats_time
