package main


import "fmt"

func main() {
	var name string
	name = "Hello World"
	fmt.Println(Hello(name))
}

func Hello(name string) string {
	return name
}
