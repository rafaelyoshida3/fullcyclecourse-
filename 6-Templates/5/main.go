package main

import (
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"content.html",
		"header.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 60},
		{"Python", 20},
	})
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8282", nil)
}
