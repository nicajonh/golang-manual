package Concurrency

import (
	"fmt"
	"time"
)

type sharded_var struct{
	reader chan int
	writer chan int
}

func sharded_var_watchdog(v sharded_var){
	go func(){
		//初始值
		var value int =0
			for {
				select {
				//监听读写
				case value = <-v.writer:
					fmt.Println("value writed!")
				case v.reader <- value:
					fmt.Println("value readed!")
				}
			}
	}()
}
//避免超时阻塞
func never_leak(ch chan int){
	timeout :=make(chan bool,1)
	go func (){
		time.Sleep(1*time.Second)
		timeout<-true
	}()
	select {
		case <-ch:
		case <-timeout:
	}
}

func main() {
	//初始化，并开始维护协程
	v := sharded_var{make(chan int), make(chan int)}
	sharded_var_watchdog(v)

	//读取初始化
	fmt.Println(<-v.reader)
	//写入一个值
	v.writer <- 1
	//读取新写入的值
	fmt.Println(<-v.reader)
}
