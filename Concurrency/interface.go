package Concurrency

import (
	"errors"
	"fmt"
)

var (
	ERR_ELEM_EXISTS=errors.New("elements exits")
	ERR_ELEM_NOT_EXISTS=errors.New("element not exits")
)


//定义切片,支持interface{}类型，相当于泛型
//假设slice不支持存储相同元素
type SomeSlice []interface{}

//初始化slice实例
func NewSomeSlice() SomeSlice{
	return make(SomeSlice,0)
}
// 定义一个用于对比的接口
type Comparable interface {
	IsEqual(a interface{}) bool
}

//定义struct类型
type Employee struct{
	Id int32
	Name string
}

//Employee类型实现了Comparable接口
func(self *Employee) IsEqual(b interface{}) bool{
	if em,ok:=b.(Employee);ok{
		return self.Id==em.Id
	}else {
		return false
	}
}

//isEqual函数用于种类型之间比较
func isEqual(a interface{},b interface{}) bool{
	if cmpa,ok :=a.(Comparable);ok{
		return cmpa.IsEqual(b)
	}else if cmpb,ok := a.(Comparable);ok{
		return cmpb.IsEqual(a)
	}else {
		return a==b
	}
}

//从slice中删除元素
func (self *SomeSlice) Remove(elem interface{}) error{
	for k,v:=range *self{
		if isEqual(v,elem){
			if k==len(*self)-1{
				*self=(*self)[:k]
			}else {
				*self=append((*self)[:k],(*self)[k+1:]...)
			}
			return nil
		}
	}
	fmt.Printf("[Error]No such element: %v\n", elem)
	return ERR_ELEM_NOT_EXISTS
}

//从slice添加元素
func(self *SomeSlice) Add(elem interface{}) error{
	for _,v:=range *self{
		if isEqual(v,elem){
			fmt.Printf("[Error]Cannot add the same element: %v\n", elem)
			return ERR_ELEM_EXISTS
		}
	}
	*self=append(*self,elem)
	return nil
}

func main() {
	// 初始化slice
	slice := NewSomeSlice()

	// 正常情况下添加不同类型元素
	slice.Add(5)
	slice.Add("huahua")
	slice.Add(Employee{Id: 123, Name: "xiaohong"})
	slice.Add(10)
	slice.Add("xiaoming")
	slice.Add(Employee{Id: 456, Name: "xiaogang"})
	fmt.Println("After Add, Current Slice:", slice)

	// 添加了重复的元素
	slice.Add(10)
	slice.Add("huahua")
	slice.Add(Employee{Id: 456, Name: "xiaogang"})
	fmt.Println("After invalid Add, Current Slice:", slice)

	// 正常情况下删除元素
	slice.Remove(5)
	slice.Remove("huahua")
	slice.Remove(Employee{Id: 456, Name: "xiaogang"})
	fmt.Println("After Remove, Current Slice:", slice)

	// 删除并不存在的元素
	slice.Remove(11)
	slice.Remove("somename")
	slice.Remove(Employee{Id: 789, Name: "dajiu"})
	fmt.Println("After invalid Remove, Current Slice:", slice)
}
