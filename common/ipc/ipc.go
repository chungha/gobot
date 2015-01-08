package ipc

import (
//"fmt"
)

//type IpcRequest struct {
//	Cmd  string
//	Strs []string
//}

//type IpcReply struct {
//	Strs []string
//}

type Ipc struct {
	//AdminMsgs  chan string
	//SearchMsgs chan string
	CallBackIntf *Ipc_func_tmpl
}

func (ii *Ipc) Req(key *IpcRequest, value *IpcReply) error {
	//n := len(key.Strs)

	tmp := *ii.CallBackIntf
	if tmp != nil {
		return tmp.Req(key, value)
	}

	////	send message...
	//if ii.AdminMsgs != nil {
	//	if key.Cmd == "admin" {
	//		ii.sendMsgs(n, ii.AdminMsgs, key.Strs)
	//	}
	//} else {
	//	fmt.Println("msgs channel is not exist.")
	//}

	//if ii.SearchMsgs != nil {
	//	if key.Cmd == "search" {
	//		ii.sendMsgs(n, ii.SearchMsgs, key.Strs)
	//	}
	//} else {
	//	fmt.Println("msgs channel is not exist.")
	//}

	return nil
}

//func (ii *Ipc) sendMsgs(n int, ch chan string, strs []string) {
//	for i := 0; i < n; i++ {
//		ch <- strs[i]
//	}
//	ch <- ""
//}
