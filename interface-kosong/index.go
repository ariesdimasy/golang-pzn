package main

import "fmt"

func main() {
	fmt.Println(hello("Dimas"))
}

func hello(name string) interface{} {
	return "Iam " + name + ", anything i want to be"
}
