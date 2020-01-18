package Concurrency

import (
	"fmt"
)

func doSomething(i int ,xi int){
	fmt.Println("i=%d,xi=%d",i,xi)
}

func main(){
	data := []int{1, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3, 2, 3, 1, 2, 3}
	N :=len(data)
	sem :=make(chan int,N)
	for i,xi := range data{
		go func(i int ,xi int){
			doSomething(i,xi)
			sem<-0
		}(i,xi)
	}
	//等待结束，查看跑完结果
	for i := 0; i < N; i++ {
		<-sem
	}
	fmt.Println("搞定了。")
}
