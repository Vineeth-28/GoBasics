package main

import (
	"fmt"
	"slices"
)

// slices = dynamic arrays
// most used construct in go
//+ useful methods .
func main(){
     
	//uninitiated slice is nil
	var nums []int      // var <slice name> [dont mention size] datatype
	fmt.Println(nums)
	fmt.Println(len(nums))


	var nums2 =make([]int ,2 , 4) //normally we put  [0,4 ] so the slice becomes empty with 0 values and capacity 4
	fmt.Println(nums2)
	//capacity
	fmt.Println(cap(nums2))


	//inserting value 
	nums2 =append(nums2,2 )
	nums2 =append(nums2,3 )
	nums2 =append(nums2,3 )
	nums2 =append(nums2,2 )

	fmt.Println(nums2)
	fmt.Println(cap(nums2))


	//copy function

	var nums3 =make([]int ,0,5)
	var nums4 =make([]int ,len(nums3))
	nums4= append(nums3, 2 )
	copy(nums4,nums3)
	fmt.Println(nums4)


	//slice operator 
	var nums5= []int {1,2,3,1,3}
	fmt.Print(nums5[0:4])

	fmt.Println(nums5[:1])


	//comparing
	var num6  = []int {1,2}
	var nums7 = []int{1,2,3}
	fmt.Println(slices.Equal(num6,nums7))

	//2d slices
      var nums8 = [][] int {{1,2},{3,4}}
	  fmt.Println(nums8)
	  
}
