package main

import "fmt"

func counter() func() int {
	var counter int = 0
	return func() int {
		counter += 1
		return counter
	}
}
func main() {
	increament := counter()
	fmt.Println(increament())
	fmt.Println(increament())
}
