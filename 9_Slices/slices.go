package main

import "fmt"

//slices = dynamic arrays
// most used construct in go
// + useful methods .
func main(){
     
	//uninitiated slice is nil
	var nums []int      // var <slice name> [dont mention size] datatype
	fmt.Println(nums)
	fmt.Println(len(nums))


	var nums2 =make([]int ,2 , 4)
	fmt.Println(nums2)
	//capacity
	fmt.Println(cap(nums2))


	//inserting value 
	nums2 =append(nums2,2 )
	nums2 =append(nums2,3 )
	nums2 =append(nums2,3 )
	nums2 =append(nums2,2 )

	fmt.Println(nums2)
}
