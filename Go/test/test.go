package main

import "fmt"

var name string = "Abdullokhon"

func main() {
	test(name)
}

func test(name string) {
	fmt.Println("Hello " + name + "!")
}
