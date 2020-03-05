package services

import (
	"bufio"
	"log"
	"os"
	"text/template"
)

func RenderYml(data interface{}) {
	f, err := os.Create("prom.yaml")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)

	tmpl := template.Must(template.ParseFiles(os.Getenv("PROMTPL")))

	tmpl.Execute(w, data)

	w.Flush()
	defer f.Close()
}
