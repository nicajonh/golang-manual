package SimpleRPC

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct{
	Name string
	Args []interface{}
}


func encode(data RPCData)([]byte,error){
	var buf bytes.Buffer
	buffEnc := gob.NewEncoder(&buf)
	if err:=buffEnc.Encode(data);err!=nil{
		return nil,err
	}
	return buf.Bytes(),nil
}

func decode(byts []byte) (RPCData,error){
	buf:=bytes.NewBuffer(byts)
	bufDec:=gob.NewDecoder(buf)
	var data RPCData
	if err:=bufDec.Decode(&data);err!=nil{
		return data,err
	}
	return data,nil
}
