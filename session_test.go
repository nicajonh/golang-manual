package main

import (
	"net"
	"sync"
	"testing"
	sRpc "./SimpleRPC"
)

func TestSession_ReadWrite(t *testing.T){
	addr := "0.0.0.0:2333"
	cont := "yep"
	wg :=sync.WaitGroup{}
	wg.Add(2)
	//Write
	go func() {
		defer wg.Done()
		l, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := l.Accept()
		s := sRpc.NewRpcSession(conn)
		err = s.Write([]byte(cont))
		if err != nil {
			t.Fatal(err)
		}
	}()
	//Read
	go func(){
		defer wg.Done()
		conn,error:=net.Dial("tcp",addr)
		if error!=nil{
			t.Fatal(error)
		}
		s:=sRpc.NewRpcSession(conn)
		data,error:=s.Read()
		if error!=nil{
			t.Fatal(error)
		}
		if string(data)!=cont{
			t.FailNow()
		}
	}()
	wg.Wait()
}