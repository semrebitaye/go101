package main

import "fmt"

func main() {
	age := 27
	const isCool = true
	name, email := "semre", "semre@gmail.com"
	size := 1.68

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T", size)
}
