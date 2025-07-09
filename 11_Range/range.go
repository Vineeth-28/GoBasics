package main

import "fmt"

//iteration over data structures
func main (){
 
	nums:=[]int {6,7,3}
	for i:=0; i<len(nums) ; i++{
	fmt.Println(nums)
	}

	m:=map[string]string{"name":"john" ,"lastname":"doe"}
	fmt.Println(m)

	//iterating on m
	for key, val :=range m{
		fmt.Println(key, val)
	}

	 //unicode code point rune
	for i,c :=range "golang"{
		fmt.Print(i,c)
	}

}