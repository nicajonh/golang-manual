package Concurrency

import (
	"math/rand"
	"fmt"
)

func rand_generator_1() int  {
	return rand.Int()

}

func main(){
	fmt.Println(rand_generator_1())
}