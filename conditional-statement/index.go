package main

import "fmt"

func main() {
	name := "dimas"
	age := 32

	if name != "" {
		fmt.Println("Hello Dimas")
	}

	if age > 0 {
		fmt.Println("you are ", age, " years old")
	}
}
