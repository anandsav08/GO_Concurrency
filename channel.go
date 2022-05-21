package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)

	// Channel : When number of messages are known in advance

	go func(){
		for i:=0;i<4;i++{
			fmt.Printf("[1] Sending %d\n",i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i:=0;i<4;i++{
		val := <- ch
		fmt.Printf("[1] Received %d\n",val)
	}
	// Channel auto close method, when number of messages aren't known
	go func(){
		for i:=0;i<3;i++{
			fmt.Printf("Sending %d\n",i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	
	for i:=range ch{
		fmt.Printf("Received %d\n",i)
	}
}