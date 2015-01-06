package ipcadapt

import (
	"fmt"
	"gobot/common/ipc"
	"log"
	"net/rpc"
)

func CallIPC(cmd string, v *string) []string {
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	list := make([]string, 0)
	list = append(list, *v)

	req := &ipc.IpcRequest{cmd, list}
	ret := &ipc.IpcReply{}
	err = client.Call("Ipc.Req", &req, &ret)
	if err != nil {
		log.Fatal("ipc error:", err)
	}

	fmt.Printf("callIPC: %v, %v", req.Strs, ret.Strs)

	if len(ret.Strs) > 0 {
		return ret.Strs
	}

	return []string{}
}
