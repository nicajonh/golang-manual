package main

import (
	"fmt"
	"reflect"
)

type Injector struct{
	mappers map[reflect.Type]reflect.Value // 根据类型map实际的值
}
var inj *Injector

func (inj *Injector) SetMap(value interface{}){
	inj.mappers[reflect.TypeOf(value)]=reflect.ValueOf(value)
}

func (inj *Injector) GetMap(t reflect.Type) reflect.Value{
	return inj.mappers[t]
}

func (inj *Injector) Invoke(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}
	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValues[k] = inj.GetMap(t.In(k))
	}
	ret := reflect.ValueOf(i).Call(inValues)
	return ret
}

func Host(name string, f func(a int, b string) string) {
	fmt.Println("Enter Host:", name)
	fmt.Println(inj.Invoke(f))
	fmt.Println("Exit Host:", name)
}

func Dependency(a int, b string) string {
	fmt.Println("Dependency: ", a, b)
	return `injection function exec finished ...`
}

func main(){
	//创建注入器
	inj=&Injector{make(map[reflect.Type]reflect.Value)}
	inj.SetMap(3030)
	inj.SetMap("zdd")

	d:=Dependency
	Host("zddhub",d)
	inj.SetMap(8080)
	inj.SetMap("www.zddhub.com")
	Host("website", d)
}