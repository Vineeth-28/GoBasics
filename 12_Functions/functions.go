package main

import "fmt"

func add (a int  ,b int ) int { //syntax  func function name (variable datatype , variable datatype2) return  datatype value {}
	return a +b

}

//same parameters -if the parameters is same it can be declared like this 
//syntax  func <name of function > (variable 1 , variable 2  datatype ) return type{}
 func add2(a,b int) int {
	return a+b
 }
func main(){
   fmt.Println( add(3,5))
   fmt.Println(add2(2,3))
}
