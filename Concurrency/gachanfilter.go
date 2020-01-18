package Concurrency

import (
	"fmt"
)

//ch是只读chan
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		//fmt.Println("Generate:i:", i)
		ch <- i
	}
}

func Filter(in <-chan int,out chan<-int,prime int){
	for {
		i:=<-in
		if i%prime!=0{
			out<-i
		}
	}
}

func main(){
	var numbers []int
	ch := make(chan int)
	go Generate(ch)
	for i:=0;i<1000;i++{
		prime:=<-ch
		numbers=append(numbers,prime)
		fmt.Println(prime, "n")
		ch1:=make(chan int)
		Filter(ch,ch1,prime)
		ch=ch1
	}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}
