package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	firstHandler := func(w http.ResponseWriter, r *http.Request) {
		// built-in template library is safe against code injection
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		// films are being passed as data, w will write data to the screen
		err = tmpl.Execute(w, films)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	secondHandler := func(w http.ResponseWriter, r *http.Request) {
		// thanks to that method we can check if request comes from HTMX or not
		// log.Print(r.Header.Get("HX-Request"))
		// key must be the same as value of the name attribute inside input element/tag in HTML/HTMX
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl, _ := template.ParseFiles("index.html")
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", firstHandler)
	http.HandleFunc("/add-film", secondHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
