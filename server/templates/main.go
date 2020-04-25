package main

import (
	"net/http"
	"text/template"
)

// Todo -
type Todo struct {
	Task      string
	Completed bool
}

func main() {
	todos := []Todo{
		Todo{Task: "task one", Completed: true},
		Todo{Task: "task two", Completed: true},
		Todo{Task: "task three", Completed: false},
		Todo{Task: "task four", Completed: true},
		Todo{Task: "task five", Completed: false},
		Todo{Task: "task six", Completed: true},
	}

	tmpl := template.Must(template.ParseFiles("template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ Todos []Todo }{
			Todos: todos,
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
