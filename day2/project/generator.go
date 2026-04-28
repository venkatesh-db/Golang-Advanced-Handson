
package main

import (
	"os"
	"text/template"
)

type Method struct {
	Name          string
	Params        string
	Return        string
	DefaultReturn string
}

type TemplateData struct {
	Name    string
	Methods []Method
}

func main() {

	data := TemplateData{
		Name: "UserService",
		Methods: []Method{
			{
				Name:          "GetUser",
				Params:        "id int",
				Return:        "string",
				DefaultReturn: `"mock-user"`,
			},
			{
				Name:          "CreateUser",
				Params:        "name string",
				Return:        "error",
				DefaultReturn: "nil",
			},
		},
	}

	tmpl := template.Must(template.ParseFiles("template.tmpl"))

	file, err := os.Create("user_mock.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl.Execute(file, data)
}