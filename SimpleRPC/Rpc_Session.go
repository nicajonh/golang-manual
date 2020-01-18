package SimpleRPC

import (
	"encoding/binary"
	"io"
	"net"
)

type Session struct{
	conn net.Conn
}

//write date to conn
func (self *Session)  Write(data []byte) error{
	buf:=make([]byte,4+len(data))// 4 字节头部 + 数据长度
	binary.BigEndian.PutUint32(buf[:4],uint32(len(data)))//大端,高位高地址
	copy(buf[4:],data)
	if _,error:=self.conn.Write(buf);error!=nil{
		return error
	}
	return nil
}
//new a session
func NewRpcSession(conn net.Conn) *Session{
	return &Session{conn:conn}
}

//read data from conn
func(self *Session) Read() ([]byte,error){
	header:=make([]byte,4)
	if _,error:=io.ReadFull(self.conn,header);error!=nil{
		return nil,error
	}
	dataLen:=binary.BigEndian.Uint32(header)
	data:=make([]byte,dataLen)
	if _,error:=io.ReadFull(self.conn,data);error!=nil{
		return nil,error
	}
	return data,nil
}

