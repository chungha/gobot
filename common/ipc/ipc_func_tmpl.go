package ipc

import (
//"fmt"
)

type IpcRequest struct {
	Cmd  string
	Strs []string
}

type IpcReply struct {
	Strs []string
}

type Ipc_func_tmpl interface {
	Req(key *IpcRequest, value *IpcReply) error
}
