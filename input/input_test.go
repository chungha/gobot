
package input

import (
  "testing"
  "fmt"

)

func Test_Input(*testing.T) {
  inputs := Input("urls");

  for e := inputs.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }
  fmt.Println("total=", inputs.Len())
}

