package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Emails []string
	Phone  string
}

const templ = `{{$name := .Name}}
{{$phone := .Phone}}
{{range .Emails}}
		{{$email := .}}
		Name is {{$name}}, email is {{$email}}, {{$phone}}
{{end}}
`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"1@gmail.com", "2@gmail.com"},
		Phone:  "010-555-5555",
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
