package reducer

import (
  "testing"
  "container/list"
)

func TestReduce(t *testing.T) {
  l := list.New()
  l.PushBack("word1 url_1 3")

  Reduce(l)

  rmap := Query("word1")
  _, ok := rmap["url_1"]

  if !ok {
    t.Error("nope url_1")
  }
}
