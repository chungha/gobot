package indexer

import (
  "testing"
  "fmt"
)

func Test_Split(*testing.T) {
  l := Split("<xxx>I am a boy.</xxx>")
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      fmt.Printf("%s\n", v);
    }
  }
}

func Test_Tokenize(*testing.T) {
  words := Tokenize("I am a Boy.");
  fmt.Printf("%q\n", words);
}
