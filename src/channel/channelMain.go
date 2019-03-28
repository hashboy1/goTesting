package main

import (
	"fmt"
	//"time"
	"strconv"
)


func add(a int,b int,pipe chan int){
	var sum int
	sum = a + b
	
	pipe <- sum

}

func main(){ 

	var pipe chan int //"chan int" is to defined the pipeline for integer type
	pipe = make(chan int,10)
	for i:=0;i<10;i++{
	go add(i,i+1,pipe)
	var result int
	result = <-pipe
	fmt.Println("result:" + strconv.Itoa(result))
	}
	//time.Sleep(5 * time.Second)
	

}