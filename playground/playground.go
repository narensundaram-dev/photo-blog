package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

func main() {
	user := struct {
		Username string
		Password string
	}{"Naren", "Sundaram"}

	v := reflect.ValueOf(user)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%v: %v\n", t.Field(i).Name, v.Field(i))
	}

}
