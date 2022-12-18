package main

import "net/http"

func main() {

	// cria um mux para ser passado ao handlefunc
	// é iniciado um web server na prota 8080 "home"
	// que escreve "Hello World"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	//mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

// é criada uma struct vaia (ou não) apenas para
// que seja implementado na função ServeHTTP
// funcionaria da mesma maneira se fosse vazia, ao
//  ser chamado o objeto em mux.Handle ele executaria
// a função ServeHTTP mesmo com "blog{}"
type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
