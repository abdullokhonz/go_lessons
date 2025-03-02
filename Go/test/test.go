package main

import "fmt"

const name string = "Abdullokhon"
var age int = 16

func main() {
	test(name, age)
}

func test(name string, age int) {
	fmt.Println(fmt.Sprintf("Hello %s(%d)!", name, age))
}
