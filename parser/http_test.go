package parser

import (
  "testing"
  "fmt"
  "strings"
)

func Test_Download(*testing.T) {
  html, err := Download("http://daum.net")
  if (err == nil) {
    fmt.Printf(html)
  } else {
    fmt.Printf("[ERROR] %s\n", err);
  }
}

func Test_Split(*testing.T) {

  words := Split("<h1>0</h1><a>1</a>2<input type=\"text\" />3<hr/>4<br/>5<p>some sentence with white space.</p>")
  for index, word := range words {
    fmt.Printf("%d. (%d) %s\n", index, len(word), word);
	if (len(word) > 0) {
	  ts := strings.Fields(word)
      fmt.Println(ts, len(ts)) // [one two three four] 4
	}
  }
}