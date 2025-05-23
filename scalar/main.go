package main

import (
	_ "embed"
	"flag"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//go:embed scalar.html
var index string

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	openapi := flag.String("openapi", "schema.json", "API definition file")
	gateway := flag.String("proxy", "", "Proxy address")
	auth := flag.String("auth", "", "Prefill authentication")
	flag.Parse()

	t := template.Must(template.New("index").Parse(string(index)))
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, template.JS(*auth))
	})

	http.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *openapi)
	})

	if *gateway != "" {
		target, err := url.Parse("http://gateway:8080")
		if err != nil {
			log.Fatal(err)
		}
		http.HandleFunc("/", httputil.NewSingleHostReverseProxy(target).ServeHTTP)
	}

	http.ListenAndServe(*addr, nil)
}
