package main

import (
	"gobot/backend_indexer/ipcadapt"
	"gobot/common/ipc"
)

func main() {
	done := make(chan bool)

	cb := new(ipc.Ipc)

	a := ipc.Ipc_func_tmpl(new(ipcadapt.Ipc_callback))
	cb.CallBackIntf = &a

	ipcadapt.Start(cb)

	<-done
}
