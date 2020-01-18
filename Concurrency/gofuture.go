package Concurrency

import (
	"fmt"
)

type query struct{
	sql chan string
	result chan string
}



func execQuery(q query){
	go func(){
		sql:=<-q.sql
		q.result <-"get"+sql
	}()
}


func main(){
	q:=query{make(chan string,1),make(chan string,1)}
	execQuery(q)
	//准备参数
	q.sql <- "select * from table"
	//获取结果
	fmt.Println(<-q.result)
}