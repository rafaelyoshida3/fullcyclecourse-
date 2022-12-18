package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// cria um http client
	c := http.Client{}

	// importante passo: bufferiza a string de bytes {"name": "wesley"}
	// passo necessário para que possa ser usado no método Post.
	// do contrário (sem bufferizar primeiro) não é possível
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))

	// aplica o método post usando do novo buffer (jsonVar) criado
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// copia o buffer de resp.Body e joga em os.Stdout
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
