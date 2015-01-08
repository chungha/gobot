package searchpage

import (
	"fmt"
	"gobot/frontend_server/ipcadapt"
	"net/http"
)

func MakePage(w http.ResponseWriter, r *http.Request) {

	list := []string{}

	v := r.URL.Query()
	for a, b := range v {
		fmt.Println(a, b)
		//	send url address to backend
		if len(b) > 0 {
			list = ipcadapt.CallIPC("search", &b[0])
		}
	}

	page := MakeHtmlPage(list)
	fmt.Fprintf(w, page)
}
