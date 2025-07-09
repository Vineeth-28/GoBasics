package main

import "fmt"

//variadic functions are the functions in which you pass the n no of values 

func sum (nums ... int )int {
	total:=0
	for _, num := range nums {
		total = total + num
	}
	return total
}

//if you want to send anything in variadic functions use syntax like this  func <name of function> (nums .. interface{}) int {}. (Interface is like any type in go language )
func main(){
  
	nums2:= []int {2,3,1,3}
	results:= sum(nums2...) //if you want to send slices you can do like this 
	fmt.Println(1,2,3,4,1,"hello",true)
	fmt .Println(sum(1,3,1,4,2,1))
	fmt.Println(results);

}