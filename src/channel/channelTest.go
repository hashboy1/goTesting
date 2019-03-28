package main

import "fmt"
func channelTest(){ 
		pipe:=make(chan int,3)
		pipe<-1
		pipe<-2
		pipe<-3
		var content int
		content = <-pipe
		fmt.Println(content)
		content = <-pipe
		fmt.Println(content)
		content = <-pipe
		fmt.Println(content)

}