package main

import "fmt"

//function for slicing int
func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

//Function for slicing sting 
func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

//generics with any 

func printSliceGeneric[T any] (items []T) { //passing any is not a good practise , you can use interface too
	for _, item := range items {
		fmt.Println(item)
	}
}

//Compareable
func printSliceGenericComperableMethod[T  comparable] (items []T) { //passing any is not a good practise , you can use interface too
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSliceGenericWithChoice[T int | string ] (items []T) { //passing any is not a good practise , you can use interface too
	for _, item := range items {
		fmt.Println(item)
	}
}


func main() {
  printSlice([]int {1,2,3,1})
  printStringSlice([]string {"golang","js","typescript"})
  printSliceGeneric([]string {"s","s","s" })
  printSliceGeneric([]bool {true ,false })
  printSliceGenericWithChoice([]int {1,2,3,1})
  printSliceGenericComperableMethod([]int {1,2,1,23})

}

// make me a bun revision runtime notes for quick revision when running into interview make a markdown explore the Internet bun documentation and everywhere for the important topics and add the bullet points examples for it