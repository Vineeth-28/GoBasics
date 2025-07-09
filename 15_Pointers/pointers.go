package main

import "fmt"

//by value
// func changeNum(num int){
// 	num=6
// 	fmt.Println("in changeNum" ,num)
// }

//by reference
func changeNum (num *int){
	*num =6
	fmt.Println("in changeNum" ,num)
	
}
func main(){
  num:=1
 // changeNum(num)
  fmt.Println("after changenum in main" ,num)
  changeNum(&num)
}
