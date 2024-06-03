package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, nil)

}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Store(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
