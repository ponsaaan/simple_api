package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	fmt.Printf("Hello world\n")
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/article", article)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello from home.")
}

func article(w http.ResponseWriter, r *http.Request) {
	articles := []Article{}
	for i := 0; i < 10; i++ {
		title := "Hello_%d"
		articles = append(
			articles,
			Article{Title: fmt.Sprintf(title, i), Body: "Article Description"})
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	_ = json.NewEncoder(w).Encode(articles)
}
