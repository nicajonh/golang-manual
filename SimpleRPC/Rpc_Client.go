package SimpleRPC

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

//new a client
func NewRpcClient(conn net.Conn)  *Client{
	return &Client{conn:conn}
}

func (c *Client) CallRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()

	// 完成与 Server 的交互
	f := func(args []reflect.Value) []reflect.Value {
		// 处理输入参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}

		// 编码 RPC 数据并请求
		cliSession := NewRpcSession(c.conn)
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, _ := encode(reqRPC)
		_ = cliSession.Write(b)

		// 解码响应数据，得到返回参数
		respBytes, _ := cliSession.Read()
		respRPC, _ := decode(respBytes)

		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行 nil 转换
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}