package indexer

import (
  "testing"
  "fmt"
)

func Test_Split(t *testing.T) {
  l := Split("<xxx>I am a boy.</xxx>")
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      fmt.Printf("%s\n", v);
    }
  }
}

func Test_Tokenize(t *testing.T) {
  words := Tokenize("I am a Boy!.");
  fmt.Printf("%q\n", words);

  if words[3] != "Boy" {
    t.Error("The third result. Expected : Boy, but " + words[3])
  }
}

func Test_Indexing(t *testing.T) {
  var indexMap = IndexingToMap("http://en.wikipedia.org/wiki/Peace");
}
