package main

import (
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	log.Println("Request", req.Method, "from", req.RemoteAddr)
	_, err := res.Write([]byte("hello"))
	if err != nil {
		log.Println("Occured error:", err.Error())
		return
	}
}

func main() {
	log.Println("Server started...")
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil )
	if err != nil {
		log.Println("Occured error:", err.Error())
	}
}
