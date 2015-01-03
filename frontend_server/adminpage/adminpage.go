package adminpage

import (
	"fmt"
	//"html"
	"net/http"
	"os"
)

func MakePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "admin, %q", html.EscapeString(r.URL.Path))

	page, _ := loadHTML("adminpage/admin.html")
	fmt.Fprintf(w, page)

	v := r.URL.Query()
	for a, b := range v {
		fmt.Println(a, b)
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
