// test_2 project main.go
package ipcadapt

import (
	"gobot/common/ipc"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//	server start
func Start(Ipc_backend *ipc.Ipc) {
	//go sample2_sub(ipc_backend)

	rpc.Register(Ipc_backend)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
