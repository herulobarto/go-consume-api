package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int64  `json:"userId"`
}

var BASE_URL = "https://jsonplaceholder.typicode.com"

func Index(w http.ResponseWriter, r *http.Request) {

	var posts []Post

	response, err := http.Get(BASE_URL + "/posts")
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&posts); err != nil {
		log.Print(err)
	}

	data := map[string]interface{}{
		"posts": posts,
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, data)

}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Store(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
