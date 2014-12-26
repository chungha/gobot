package indexmap

import (
  "fmt"
  "testing"
  "container/list"
)

func TestAdd(t *testing.T) {
  indexMap := New();

  indexMap.Add("word", "url", 0)

  posList, ok := indexMap.QueryWordAndUrl("word", "url")
  if !ok {
    t.Error("nope url")
    return
  }
  if posList.Len() != 1 {
    t.Error("rlist's length should be 1")
    return
  }
}

func TestAddIndexStringListMultiURL(t *testing.T) {
  l := list.New()
  l.PushBack("word1 url_1 3")
  l.PushBack("word1 url_1 2")
  l.PushBack("word1 url_2 3")

  indexMap := New()

  indexMap.AddIndexStringList(l)

  rmap := indexMap.QueryWord("word1")
  rlist, ok := rmap["url_1"]
  if !ok {
    t.Error("nope url_1")
    return
  }

  if rlist.Len() != 2 {
    t.Error(fmt.Sprint("Result list length should be 2. but %d", rlist.Len()));
    return
  }
  iter := rlist.Front()
  if iter.Value.(int) != 3 {
    t.Error("Result list should contain 3. but %d", rlist.Front().Value.(int));
    return
  }
  iter = iter.Next()
  if iter.Value.(int) != 2 {
    t.Error("Result list should contain 2. but %d", rlist.Front().Value.(int));
    return
  }

  rlist, ok = rmap["url_2"]
  if !ok {
    t.Error("nope url_2")
    return
  }

  if rlist.Len() != 1 {
    t.Error(fmt.Sprint("Result list length should be 1. but %d", rlist.Len()));
    return
  }
  iter = rlist.Front()
  if iter.Value.(int) != 3 {
    t.Error("Result list should contain 3. but %d", rlist.Front().Value.(int));
    return
  }
}
