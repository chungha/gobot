package reducer

import (
  "testing"
  "fmt"
  "container/list"
)

func TestInput(*testing.T) {
  l := list.New()
  l.PushBack("word1 url_1 3");
  l.PushBack("word1 url_1 1");
  l.PushBack("word1 url_1 5");
  l.PushBack("word1 url_2 3");
  l.PushBack("word1 url_2 4");
  l.PushBack("word2 url_2 5");
  Reduce(l)

  fmt.Printf(Query("word1"))
}
