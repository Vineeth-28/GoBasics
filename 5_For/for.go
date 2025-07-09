package main

import "fmt"

//for - only construct in go for looping
func main(){
  
	//while is not present and the way to write it  is below 
	i:=1
	for i<= 3{
		fmt.Println(i)
		i++
	}


	//infinite loop
	// for {
	// 	fmt.Println(1)
	// }
 
	//classic for loop
  for j:=0;  j<5; j++ {
	fmt.Println(j)
  }

  //range 
  for i:= range 4{
	fmt.Println(i)
  }

}