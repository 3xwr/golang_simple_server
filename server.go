package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/headers", headerRequest)
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	requestedFile :=  "./" + r.URL.Path
	http.ServeFile(w,r,requestedFile)
	return
}

func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	return
}

func headerRequest(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
