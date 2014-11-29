package parser

import (
  "testing"
  "fmt"
)

func Test_Split(*testing.T) {
  l := Indexing("http://daum.net")
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      fmt.Printf("%s\n", v);
    }
  }
}
