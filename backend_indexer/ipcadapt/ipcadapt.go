// test_2 project main.go
package ipcadapt

import (
	"gobot/common/ipc"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type IpcBackend struct {
	ipc_backend *ipc.Ipc
}

//	initialize resource
//	return: admin channel, search channel
func (ii *IpcBackend) Prepare() (*chan string, *chan string) {
	ii.ipc_backend = new(ipc.Ipc)

	ii.ipc_backend.AdminMsgs = make(chan string)
	ii.ipc_backend.SearchMsgs = make(chan string)

	return &ii.ipc_backend.AdminMsgs, &ii.ipc_backend.SearchMsgs
}

//	server start
func (ii *IpcBackend) Start() {
	//go sample2_sub(ipc_backend)

	go rpc.Register(ii.ipc_backend)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
