package searchpage

import (
	"fmt"
	"gobot/frontend_server/ipcadapt"
	"net/http"
)

type SearchPage struct {
	List []string
}

func (inst *SearchPage) MakePage(w http.ResponseWriter, r *http.Request) {

	//list := []string{"test haha !!!"}

	v := r.URL.Query()
	for a, b := range v {
		fmt.Println(a, b)
		//	send url address to backend
		if len(b) > 0 {
			inst.List = ipcadapt.CallIPC("search", &b[0])
		}
	}

	page := MakeHtmlPage(inst.List)
	fmt.Fprintf(w, page)
}
