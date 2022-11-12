package main

import (
	"flag"

	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"tailscale.com/tsnet"
)

var (
	origin   = flag.String("origin", "http://localhost:8080", "origin for the proxy")
	hostname = flag.String("hostname", "proxy", "hostname for the tailnet")
)

func main() {
	flag.Parse()

	s := &tsnet.Server{
		Hostname: *hostname,
	}

	defer s.Close()

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	origin, _ := url.Parse(*origin)

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.Host = origin.Host
		req.URL.Host = origin.Host
		req.URL.Scheme = origin.Scheme
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.Serve(ln, nil))
}
