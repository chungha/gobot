
package input

import (
  "testing"
  "fmt"

)

func Test_Input(*testing.T) {
  inputs := Input();

  for e := inputs.Front(); e != nil; e = e.Next() {
       fmt.Println(e.Value)
  }
}

