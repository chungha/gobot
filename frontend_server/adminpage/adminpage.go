package adminpage

import (
	"fmt"
	"gobot/frontend_server/ipcadapt"
	"net/http"
	"os"
)

func MakePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "admin, %q", html.EscapeString(r.URL.Path))

	page, _ := loadHTML("adminpage/admin.html")
	fmt.Fprintf(w, page)

	v3 := r.PostFormValue("url")
	if len(v3) > 0 {
		//	send url address to backend
		fmt.Println(v3)
		ipcadapt.CallIPC("admin", &v3)
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
