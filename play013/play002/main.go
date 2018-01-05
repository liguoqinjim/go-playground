package main

import (
	"log"
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error:%v", err)
		}
	}()

	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t := template.Must(template.New("todos").Parse("You have task named \"{{ {.Name}}\" with description: \"{{ .Description}}\""))

	err := t.Execute(os.Stdout, td)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
