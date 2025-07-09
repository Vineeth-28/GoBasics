package main

import "fmt"

func main() {
 
	 //if loop 
	age := 19
	if age >= 19 {
		fmt.Println("person is adult ")
	}
   

	// if else 
	age2 := 16
	if age2 >= 19 {
		fmt.Println("person is adult ")
	}else {
		fmt.Println( "person is not an adult")
	}

	//if elseif else 
	age3 := 22
	if age3 >= 19 {
		fmt.Println("person is adult ")
	}else if age3 <=19 {
		fmt.Println( "person is not an adult")
	}else {
		 fmt.Println("person is kid ")
	}


   //example 
	var role ="admin"
	var hasPermission =true 
	if role == "admin" || hasPermission {
		fmt.Println("yes valid")
	}

	//you can declare value in if 
    //(we can declare a variable inside if construct)
	if 	age:=15 ; age >=18 {
		 fmt.Println("person is an adult")
	}else if age >=12{
		fmt.Println("person is teenage")
	}

}