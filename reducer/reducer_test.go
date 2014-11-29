package reducer

import (
  "fmt"
  "testing"
  "container/list"
)

func TestReduceMultiURL(t *testing.T) {
  l := list.New()
  l.PushBack("word1 url_1 3")
  l.PushBack("word1 url_1 2")
  l.PushBack("word1 url_2 3")

  Reduce(l)

  rmap := Query("word1")
  rlist, ok := rmap["url_1"]
  if !ok {
    t.Error("nope url_1")
  }

  if rlist.Len() != 2 {
    t.Error(fmt.Sprint("Result list length should be 2. but %d", rlist.Len()));
  }
  iter := rlist.Front()
  if iter.Value.(string) != "3" {
    t.Error("Result list should contain 3. but " + rlist.Front().Value.(string));
  }
  iter = iter.Next()
  if iter.Value.(string) != "2" {
    t.Error("Result list should contain 2. but " + rlist.Front().Value.(string));
  }

  rlist, ok = rmap["url_2"]
  if !ok {
    t.Error("nope url_2")
  }

  if rlist.Len() != 1 {
    t.Error(fmt.Sprint("Result list length should be 1. but %d", rlist.Len()));
  }
  iter = rlist.Front()
  if iter.Value.(string) != "3" {
    t.Error("Result list should contain 3. but " + rlist.Front().Value.(string));
  }
}
