package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/skinass/go-sum"
)

func main() {
	counter := sum.NewAdder()

	http.HandleFunc("/whats_time", func(w http.ResponseWriter, r *http.Request) {
		counter.Add(1)
		w.Write([]byte(time.Now().String()))
	})

	go func() {
		for _ = range time.NewTicker(time.Second).C {
			fmt.Printf("wow! there is already %d hits!!!\n", counter.Sum())
		}
	}()

	http.ListenAndServe(":80", nil)
}

// nc -l localhost 8081

// https://github.com/aquasecurity/chain-bench/blob/main/docs/CIS-Software-Supply-Chain-Security-Guide-v1.0.pdf
