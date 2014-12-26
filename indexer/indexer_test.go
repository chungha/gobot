package indexer

import (
  "testing"
  "fmt"
)

func Test_Split(t *testing.T) {
  l := Split("<xxx>000\tAAA\r\n888<xxx>")

  if l.Len() != 3 {
    t.Error("splitted length should be 3.")
    return
  }
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      fmt.Printf("---%s---\n", v);
    }
  }
}

func Test_Tokenize(t *testing.T) {
  words := Tokenize("I am a Boy!.");
  fmt.Printf("%q\n", words);

  if len(words) >= 4 &&  words[3] != "Boy" {
    t.Error("The third result. Expected : Boy, but " + words[3])
  }
}

func Test_Indexing(t *testing.T) {
}
