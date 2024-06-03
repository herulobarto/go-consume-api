package main

import (
	"log"
	"net/http"

	post "github.com/herulobarto/go-consume-api/controllers"
)

func main() {

	http.HandleFunc("/", post.Index)
	http.HandleFunc("/posts", post.Index)
	http.HandleFunc("/post/create", post.Create)
	http.HandleFunc("/post/store", post.Store)
	http.HandleFunc("/post/delete", post.Delete)

	log.Print("Server started on port: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
