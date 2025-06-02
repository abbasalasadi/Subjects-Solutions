package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.CamelToSnakeCase("HelloWorld"))
	fmt.Println(p.CamelToSnakeCase("helloWorld"))
	fmt.Println(p.CamelToSnakeCase("camelCase"))
	fmt.Println(p.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(p.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(p.CamelToSnakeCase("hey2"))
}
