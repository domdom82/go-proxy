package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	target := flag.String("target", "", "The target URL to forward to.")

	flag.Parse()

	if *target == "" {
		panic("Missing target URL.")
	}

	targetURL, err := url.Parse(*target)

	if err != nil {
		panic("Expected URL but got: " + err.Error())
	}
	rp := httputil.NewSingleHostReverseProxy(targetURL)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = targetURL.Host
		rp.ServeHTTP(w, r)
	})
	_ = http.ListenAndServe(":8080", mux)
}
