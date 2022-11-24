package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Qadriddin",
		Age:  21,
		Meta: UserMeta{
			Visits: 5,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
