
package input

import (
  "fmt"
  "container/list"
)


func Input() *list.List {
  inputs := list.New();

  inputs.PushBack("http://en.wikipedia.org/wiki/Go_(programming_language)");
  inputs.PushBack("http://en.wikipedia.org/wiki/Peace");
  inputs.PushBack("http://en.wikipedia.org/wiki/Apple_Inc.");

  for e := inputs.Front(); e != nil; e = e.Next() {
       fmt.Println(e.Value)
  }

  return inputs
}

