package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
func main() {
	var arr = []interface{}{true, 1, "bool", func() {}, struct{}{}}
	for _, v := range arr {
		v := reflect.ValueOf(v)
		fmt.Printf("'%v' is '%s'\n", v, v.Kind().String())
		fmt.Printf("")
	}
}
