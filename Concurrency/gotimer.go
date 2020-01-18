package Concurrency

import (
	"time"
	"fmt"
)

func main()  {
	timer :=time.NewTimer(3*time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer has expired")
	}()

	//timer.Stop()
	timer.Reset(3*time.Second)
	time.Sleep(60*time.Second)
}
