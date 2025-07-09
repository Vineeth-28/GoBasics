package main

import "fmt"
import "time"

func main (){
	//simple switch
	i:=1
  switch i{
  case 1:
	fmt.Println("one ")
	
  case 2:
	fmt.Println("two")

  case 3: 
  fmt.Println("three")
  }

  //multiple condition switch

  switch time.Now().Weekday(){
     case time.Saturday ,time.Sunday:
	fmt.Println("its weekend")

	default :
	fmt.Println("its work day")

  }

  //type switch 

whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(55)
}
