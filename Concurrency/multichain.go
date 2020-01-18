package Concurrency

import (
	"fmt"
	"math/rand"
)

//直接返回同道channel
func rand_generator_2() chan int{
	//创建chan,不指定大小时,chan会读空阻塞
	out := make(chan int)
	go func() {
		for {
			//向同道内写数据,如果无人读取会阻塞
			out <- rand.Int()
		}
	}()
	return out
}

/**
* 多路复用，高并发版生成器
* 根据已知权限使用函数生成相应数据，异步调用节省了大量时间。
 */
func rand_generator_3() chan int {
	//创建两个随机数生成器服务
	rand_service_handler_1 := rand_generator_2()
	rand_service_handler_2 := rand_generator_2()
	out:=make(chan int)
	go func(){
		for {
			out <- <-rand_service_handler_1

		}
	}()
	go func(){
		for {
			out<- <-rand_service_handler_2
		}
	}()
	return out
}

func main() {
	//生成随机数作为一个服务
	rand_service_handler := rand_generator_3()
	//从服务中读取随机数并打印
	fmt.Println("%dn", <-rand_service_handler)
	fmt.Println("%dn", <-rand_service_handler)
	fmt.Println("%dn", <-rand_service_handler)
}

