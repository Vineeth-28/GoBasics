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


 //multiple return statement 
 func getLanguages () (string , string){
  return "golang" ,"javascript "
 }



 //multiple datatype return statement 
 func getLanguages2() (string , string ,bool){
  return "golang" ,"javascript ", true
 }

 func processIt( fn func (a int ) int ){
      fn(2)
 }

func main(){
   fmt.Println( add(3,5))
   fmt.Println(add2(2,3))
   fmt.Println(getLanguages())
   fmt.Println(getLanguages2())
   fn:= func  (a int )  int  {
	return 2 
   }
   processIt(fn)
}

