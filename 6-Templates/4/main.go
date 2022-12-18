package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"GO", 40},
			{"Java", 60},
			{"Python", 20},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
