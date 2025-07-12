package main

import (
	"fmt"
	"sync"
)



func task(id int, wg *sync.WaitGroup) { // passing it in function means it should be a pointer of sync.WaitGroup
	defer wg.Done()
	fmt.Println("doing task", id)
}

func main (){
	var wq sync.WaitGroup
	for i :=0 ; i<=10 ; i++{
		wq.Add(1)
		go task(i ,&wq) //to make all this synchronize we use wait group
	}

	wq.Wait();

	}

  //this is to be remeber in WaitGroup
	//add 
	//done 
	//wait


