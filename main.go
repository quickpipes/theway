package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpPort := 8080

	http.HandleFunc("/", WayServer)

	log.Printf("Starting theway server, listening on %v\n", httpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}

func WayServer(w http.ResponseWriter, r *http.Request) {

	var path string
	if r.URL.Path[1:] == "" {
		path = "this"
	} else {
		path = r.URL.Path[1:]
	}
	fmt.Fprintf(w, "%s is the way.", path)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
