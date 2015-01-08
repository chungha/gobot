package ipcadapt

import (
	"fmt"
	"gobot/common/ipc"
)

type Ipc_callback struct {
	name string
}

func (ii *Ipc_callback) Req(key *ipc.IpcRequest, value *ipc.IpcReply) error {

	fmt.Println("IPC Req: ", key.Cmd)
	switch key.Cmd {
	case "admin":
	case "search":
		testSearch(key.Strs)
		value.Strs = []string{"123", "456", "789"}
	default:
		fmt.Println("cmd error!!!")
	}

	return nil
}

//	test...
func testSearch(strs []string) {
	n := len(strs)
	//value.Strs = make([]string, n)

	for i := 0; i < n; i++ {
		//value.Strs[i] = key.Strs[i]
		fmt.Println("IPC search: ", strs[i])
	}
}
