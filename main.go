package main

import "fmt"

func main() {
	var name string
	name = "Hello World"
	Hello(name)
}

func Hello(name string) {
	fmt.Println(name)
}
