package main

import (
	"html/template"
	"os"
)

type Meta struct {
	Address string
}
type User struct {
	Name string
	Age  int
	Meta
	Bio string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Udit",
		Age:  30,
		Meta: Meta{
			Address: "1234 nangola",
		},
		Bio: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
