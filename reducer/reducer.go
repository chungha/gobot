package reducer

import (
  "fmt"
  "container/list"
)



func Reduce(l *list.List) {
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      fmt.Printf(v)
    }
  }
}
