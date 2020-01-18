package Concurrency
import (
	"fmt"
	"time"
)

func main(){
	ch :=make(chan int,1)
	timeout :=make(chan bool,1)
	//producer
	go func() {
		for i:=0;i<5;i++{
			ch<-i
			fmt.Println("生产了:%d",i)
		}
	}()

	//comsumer
	var value int
	for i :=0;i<10;i++{
		//produce a value per second
		go func() {
			time.Sleep(10 * time.Second)
			timeout <- true
		}()
		select {
		case value=<-ch:
			fmt.Println("消费了：%d", value)
		case <-timeout:
			fmt.Println("timeout!")

		}
	}
}