package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	r := mux.NewRouter()

	r.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // the title

		// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

		data := TodoPageData{
			PageTitle: fmt.Sprintf("Todo title: %s", title),
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)

	})

	http.ListenAndServe(":80", r)
}
