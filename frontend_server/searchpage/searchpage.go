package searchpage

import (
	"fmt"
	"gobot/frontend_server/ipcadapt"
	"net/http"
	"os"
)

func MakePage(w http.ResponseWriter, r *http.Request) {
	page, _ := loadHTML("searchpage/search.html")
	fmt.Fprintf(w, page)

	v := r.URL.Query()
	for a, b := range v {
		fmt.Println(a, b)
		//	send url address to backend
		if len(b) > 0 {
			ipcadapt.CallIPC("search", &b[0])
		}
	}
}

func loadHTML(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return "", err
	}

	tmpBuff := make([]byte, fi.Size())

	n, _ := file.Read(tmpBuff)
	fmt.Println("size: ", n)

	return string(tmpBuff), nil
}
