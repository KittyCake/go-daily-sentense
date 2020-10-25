package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/sentence", sentenceHandler)
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)
	fmt.Println("listening...")

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func sentenceHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

	}

	_, err = fmt.Fprintf(w, "%s", body)
}
