package reducer

import (
  "testing"
  "container/list"
)

func TestInput(*testing.T) {
  l := list.New()
  l.PushBack("aaa");
  l.PushBack("bbb");
  Reduce(l)
}
