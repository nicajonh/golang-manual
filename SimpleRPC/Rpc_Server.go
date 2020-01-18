package SimpleRPC

import (
	"fmt"
	"net"
	"reflect"
)
type Server struct{
	addr string
	funcs map[string]reflect.Value
}

//regist name of func to real method
func (s *Server) Register(rpcName string,f interface{}){
	if _,ok:=s.funcs[rpcName];!ok{
		return
	}
	fVal:=reflect.ValueOf(f)
	s.funcs[rpcName]=fVal

}
func NewRpcServer(add string) *Server{
	return &Server{add,make(map[string]reflect.Value)}
}
func (s *Server) Run(){


	//listing addr
	l,_:=net.Listen("tcp",s.addr)
	for{
		//get conn
		conn,_:=l.Accept()
		rpcSesson:=NewRpcSession(conn)
		//read data
		data,_:=rpcSesson.Read()
		//decode data
		rpcData,_:=decode(data)
		f,ok:=s.funcs[rpcData.Name]
		if !ok{
			fmt.Printf("func %s not exists", rpcData.Name)
			return
		}
		inArgs:=make([]reflect.Value,0,len(rpcData.Args))
		for _,arg:=range inArgs{
			inArgs=append(inArgs,arg)
		}
		//run it
		out:=f.Call(inArgs)
		outArgs:=make([]interface{},0,len(out))
		for _,out:=range out {
			outArgs=append(outArgs,out.Interface())
		}
		resRpcData:=RPCData{rpcData.Name,outArgs}
		//encode
		resRpcBytes,_ :=encode(resRpcData)
		_ = rpcSesson.Write(resRpcBytes)
	}

}
