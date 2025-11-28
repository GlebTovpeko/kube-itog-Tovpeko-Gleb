package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Start Info Service")

	http.HandleFunc("/info", infoHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func infoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t %s\t %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		if _, err := w.Write([]byte("Info Service")); err != nil {
			log.Println(err)
		}
	}
}
