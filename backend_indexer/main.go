package main

import (
	"fmt"
	"gobot/backend_indexer/ipcadapt"
)

func adminLoop(msgs *chan string) {
	for {
		msg := <-*msgs
		if len(msg) > 0 {
			fmt.Println("admin server: ", msg)
		}

		//	TODO: convey url to indexer..
		//GobotPmap(1)
		//GobotSequential();
	}
}

func searchLoop(msgs *chan string) {
	for {
		msg := <-*msgs
		if len(msg) > 0 {
			fmt.Println("search server: ", msg)
		}
	}
}

func main() {
	done := make(chan bool)

	reqFromUI := new(ipcadapt.IpcBackend)

	adminMsgs, seearchMsgs := reqFromUI.Prepare()

	go adminLoop(adminMsgs)
	go searchLoop(seearchMsgs)

	reqFromUI.Start()

	<-done
}
