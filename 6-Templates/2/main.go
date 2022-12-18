package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := t.Execute(os.Stdout, curso)

	//tmp := template.New("CursoTemplate")
	//tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	//err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
