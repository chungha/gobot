/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
  "gobot/input"
  "gobot/indexer"
  "gobot/reducer"
)


func main() {
  list := input.Input("input/urls")
  for e := list.Front(); e != nil; e = e.Next() {
    url := e.Value.(string)
    list := indexer.Indexing(url)
    reducer.Reduce(list)
  }
  //reducer.Print()
}
