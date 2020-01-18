package main

import (
	srpc "./SimpleRPC"
	"fmt"
	"net"
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
}
func TestRpc(t *testing.T){
	//server
	addr:="0.0.0.0:2333"
	svr:=srpc.NewRpcServer(addr)
	svr.Register("qureryUser",queryUser)
	go svr.Run()
	//client
	conn,err:=net.Dial("tcp",addr)
	if err!=nil{
		t.Error(err)
	}
	client:=srpc.NewRpcClient(conn)
	//define prototype method for remote call
	var query func(int)(User,error)
	client.CallRPC("queryUser",&query)
	u,err:=query(1)
	fmt.Println(err,u)

}

func TestMakeFunc(t *testing.T){
	swap := func(args []reflect.Value) []reflect.Value {
		return []reflect.Value{args[1], args[0]}
	}

	var intSwap func(int, int) (int, int)
	fn := reflect.ValueOf(&intSwap).Elem() // 获取 intSwap 未初始化的函数原型
	v := reflect.MakeFunc(fn.Type(), swap) // MakeFunc 使用传入的函数原型创建一个绑定 swap 的新函数
	fn.Set(v)                              // 为函数 intSwap 赋值

	fmt.Println(intSwap(1, 2)) // 2 1
}
func TestReremoteCall(t *testing.T){
	funcs := make(map[string]reflect.Value) // server 端维护 funcName => func 的 map
	funcs["incr"] = reflect.ValueOf(incr)
	args := []reflect.Value{reflect.ValueOf(1)} // 构建参数（client 传递上来）
	vals := funcs["incr"].Call(args)            // 调用执行
	var res []interface{}
	for _, val := range vals {
		res = append(res, val.Interface()) // 处理返回值
	}
	fmt.Println(res)	// [2, <nil>]
}

//define remotefunc
func queryUser(uid int) (User, error) {
	userDB := make(map[int]User)
	userDB[0] = User{"Dennis", 70}
	userDB[1] = User{"Ken", 75}
	userDB[2] = User{"Rob", 62}
	if u, ok := userDB[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("id %d not in user db", uid)
}

func incr(n int) (int, error) {
	return n + 1, nil
}