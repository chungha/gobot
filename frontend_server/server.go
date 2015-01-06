package main

import (
	"fmt"
	"gobot/frontend_server/adminpage"
	"gobot/frontend_server/searchpage"
	"log"
	"net/http"
)

func startServer() (bool, error) {
	//inst := searchpage.SearchPage{[]string{"test gara !!!"}}
	inst := searchpage.SearchPage{}
	http.HandleFunc("/search", inst.MakePage)
	http.HandleFunc("/admin", adminpage.MakePage)

	log.Fatal(http.ListenAndServe("localhost:4000", nil))

	fmt.Println("server ended...")

	return true, nil
}
