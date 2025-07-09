package main

import "fmt"

//maps = hash ,objects ,dictionary
func main (){

  //creating the map 
                            //key    //value
   m := make(map[string]string)
  //setting an elements
  m["name"] ="golang"
  m["area"] ="backend"

  //getting an element
  fmt.Println(m["name"] ,m["area"])
  fmt.Println(m["phone"]) //if key values doesnot exits then it returns zero values ==means empty 

  a:=make(map[string]int)
  a["age"] =30;
  fmt.Println("age")
  fmt.Println(a["phone"])  //if key values doesnot exits then it returns zero values ==means zero

  //length
  fmt.Print(len(m))
  fmt.Println(len(a))

  //delete function

  delete(a ,"phone")
  fmt.Println(m)

  //empty map 
  clear(m)
  fmt.Println(m)

  //map without make keyword 

   n:=map[string]int {"price ":30 ,"phone ":3} //similar to js objects 
   fmt.Println(n)


}
