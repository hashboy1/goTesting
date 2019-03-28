package main

import (
	"fmt"
	"time"
)
var pipe chan int //"chan int" is to defined the pipeline for integer type

func add(a int,b int){
	var sum int
	sum = a + b
	
	pipe <- sum

}

func main(){ 

	
	pipe :=make(chan int,1)
	go add(1,2)
	var result int
	time.Sleep(5 * time.Second)
	result = <-pipe
	fmt.Println(result)
	

}