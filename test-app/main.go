package main

import (
	"fmt"
	"reflect"

	"github.com/karincake/pentol"
)

func main() {
	got := pentol.Bool(true)
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Bool || !*got {
		fmt.Println("failed to get pointer of bool", *got)
	}
}
