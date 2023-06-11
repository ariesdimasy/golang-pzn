package main

import (
	"fmt"
)

func newmap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	fmt.Println(newmap("Dimas"))
}
