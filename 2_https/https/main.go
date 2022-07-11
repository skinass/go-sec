package main

import (
	"net/http"
	"time"

	"github.com/a-h/hsts"
)

func main() {
	http.Handle("/whats_time", hsts.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})))
	http.ListenAndServeTLS(":80", "localhost.crt", "localhost.key", nil)
}

// openssl req -new -subj "/C=RU/ST=Msk/CN=localhost" -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
// openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
//
// sudo tcpdump -A -i lo0 'port 80'
// curl -k https://localhost:80/whats_time
