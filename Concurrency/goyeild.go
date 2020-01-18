package Concurrency

import (
	"fmt"
	"math/rand"
)

/**
* 生成器
* 根据已知权限使用函数生成相应数据，异步调用节省了大量时间。
 */

//直接返回同道channel
func rand_generator_4() chan int{
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

func main(){
	rand_service_handler := rand_generator_4()
	fmt.Println("%dn", <-rand_service_handler)
	fmt.Println("%dn", <-rand_service_handler)
	fmt.Println("%dn", <-rand_service_handler)
}