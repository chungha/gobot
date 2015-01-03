package main

import (
	"fmt"
	"gobot/frontend_server/adminpage"
	"gobot/frontend_server/searchpage"
	"log"
	"net/http"
)

func startServer() (bool, error) {
	http.HandleFunc("/search", searchpage.MakePage)
	http.HandleFunc("/admin", adminpage.MakePage)

	log.Fatal(http.ListenAndServe("localhost:4000", nil))

	fmt.Println("server ended...")

	return true, nil
}
