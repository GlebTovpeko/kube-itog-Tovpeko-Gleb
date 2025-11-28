package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Start API Service")

	http.HandleFunc("/api", apiHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func apiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t %s\t %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		if _, err := w.Write([]byte("API Service")); err != nil {
			log.Println(err)
		}
	}
}
