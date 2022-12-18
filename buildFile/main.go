package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	var palavra string
	possibilidades := []string{
		"Yoshida",
		"yoshida",
		"Rafael Yoshida",
		"rafael yoshida",
		"rafael yoshida oliveira",
		"Rafael Yoshida Oliveira",
	}
	for _, input := range os.Args[1:] {
		palavra += input
	}
	println(palavra)
	z := 0

	for _, i := range possibilidades {
		if palavra == i {
			z = 1

		} else {
			z = 0
		}
	}

	if z == 1 {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Macho demais ta doido"))
		})
	} else {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hmmm tu é gayzao ne"))
		})
	}

	http.ListenAndServe(":8080", mux)
}
