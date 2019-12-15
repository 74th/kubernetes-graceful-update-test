package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	num := os.Getenv("NUM")
	sleep, err := strconv.Atoi(os.Getenv("SLEEP"))
	if err != nil {
		sleep = 3
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Print("request reserved")
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Fprintf(w, "Hello %s", num)
		log.Print("done")
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
