package main

import "fmt"

const age =30
func main(){

	 //constants are fixed value and it cannot be reassigned again 
	const name ="sf"
	fmt.Println(name);
	fmt.Println(age)


	//const:=30 (shorthand method )// this is not allowed outside the function ..

	//constant grouping (combining constants )
	const  (
		port =5500
		host="localhost"
)

	fmt.Println(port ,host)
}