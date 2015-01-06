package searchpage

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

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

func makeHtmlTable(list []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("<table border=\"1\">\r\n")

	length := len(list)
	for i := 0; i < length; i++ {
		//for str, _ := range list {
		buffer.WriteString("<tr>\r\n<td>\r\n")
		buffer.WriteString(list[i])
		buffer.WriteString("</td>\r\n</tr>\r\n")
	}
	buffer.WriteString("</table>\r\n")

	return buffer.String()
}

func MakeHtmlPage(list []string) string {
	raw, _ := loadHTML("searchpage/search.html")

	return fmt.Sprintf(strings.Replace(raw, "@@@@TABLE_REPLACE@@@@", makeHtmlTable(list), -1))
}
