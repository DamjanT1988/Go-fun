package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type Todo struct {
	Title string
	Done  bool
}

var (
	todos   []Todo
	tmpl    *template.Template
	tmplErr error
	mux     sync.Mutex
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", todos)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new.html", nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		if title != "" {
			mux.Lock()
			todos = append(todos, Todo{Title: title})
			mux.Unlock()
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/new", http.StatusMethodNotAllowed)
	}
}

func main() {
	tmpl, tmplErr = template.ParseFiles("templates/index.html", "templates/new.html")
	if tmplErr != nil {
		log.Fatalf("Error parsing templates: %v", tmplErr)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/save", saveHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
