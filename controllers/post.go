package controllers

import (
	"bytes"
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

// ada di https://jsonplaceholder.typicode.com/guide/ pilih listing all resources
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

	temp, _ := template.ParseFiles("views/create.html")
	temp.Execute(w, nil)

}

func Store(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	newPost := Post{
		Id:     0,
		Title:  r.Form.Get("post_title"),
		Body:   r.Form.Get("post_body"),
		UserId: 1,
	}

	JsonValue, _ := json.Marshal(newPost)
	buff := bytes.NewBuffer(JsonValue)

	req, err := http.NewRequest(http.MethodPost, BASE_URL+"/posts", buff)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	var postResponse Post

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&postResponse); err != nil {
		log.Print(err)
	}

	// fmt.Println(res.StatusCode)
	// fmt.Println(res.Status)
	// fmt.Println(postResponse)

	if res.StatusCode == 201 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
