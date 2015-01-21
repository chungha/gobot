package ipcadapt

import (
	"fmt"
	"gobot/backend_indexer/indexer"
	"gobot/backend_indexer/indexmap"
	"gobot/common/ipc"
)

type Ipc_callback struct {
	name     string
	indexMap *indexmap.IndexMap
}

func (ii *Ipc_callback) Req(key *ipc.IpcRequest, value *ipc.IpcReply) error {

	fmt.Println("IPC Req: ", key.Cmd)
	switch key.Cmd {
	case "admin":
		fmt.Println(key.Strs)
		ii.indexMap = GobotSequentialSingle(key.Strs[0])
	case "search":
		//testSearch(key.Strs)
		//value.Strs = []string{"123", "456", "789"}

		if ii.indexMap != nil {
			fmt.Println(key.Strs)
			tmpStr := ii.indexMap.QueryPrint(key.Strs[0])

			value.Strs = []string{tmpStr}

			fmt.Println(tmpStr)
		}

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

func GobotSequentialSingle(url string) *indexmap.IndexMap {

	var indexMap = indexmap.New()
	indexer.IndexingToMap(url, indexMap)

	return indexMap
}
