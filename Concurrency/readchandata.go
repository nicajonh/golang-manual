package Concurrency

import (
	"fmt"
	"time"
)

const(
	num = 10000000 //send data or receive data
)



func main(){
 TestChanRead()
}

func TestChanRead(){
	st :=time.Now().Unix()
	c :=make(chan int)
	go func(){
		for n:=range c{
			fmt.Printf("task TestChan2 cost %d \r\n", (time.Now().UnixNano()-st)/int64(time.Millisecond))
			fmt.Printf("read chan data:%d\n", n)
		}
	}()
	for i:=0;i<num;i++{
		c<-i
	}
	close(c)
	time.Sleep(3*time.Second)
}
