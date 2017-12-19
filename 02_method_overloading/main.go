package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func hello(name string) {
	fmt.Println("hello " + name)
}

func main() {
	hello()
}
