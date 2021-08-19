package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main(){
	PrintName()
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/success", successGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8000",nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "index.html", nil)
}
func indexPostHandler(w http.ResponseWriter, r *http.Request){
	// r.ParseForm()
	// comment := r.PostForm.Get("comment")
	http.Redirect(w, r,"/success", 302)
}

func successGetHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "success.html", nil)
}
