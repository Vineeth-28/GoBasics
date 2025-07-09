package main

import "fmt"


//arrays  => fixed size that is predictable 
                      //constant time
                     //memory optimization

func main(){
  //arrays  is sequence of specific length 
  var nums [7]int //declaration   => var <name of array  > [size]datatype
  fmt.Println(len(nums)) //len is function to check the len
 
  //inserting element in nums 
   nums[0] =2
   nums[2] =3
   nums[3] =5
   nums[5] =2
   nums[1] =2

   //getting elements 
   fmt.Println(nums[2])
   //seeing whole array
   fmt.Println(nums)

  //declaring array and inserting value at the same time
   arr:=[4]int {1,2,4,2}.   //=> name:= [size]datatype{values}
   fmt.Println(arr)


   //2dArrays
   arr2:=[2][2]int{{1,2},{3,4}}  //name:=[size][size]datatype{{  values }  ,{  values } } 
   fmt.Println(arr2)

}
