/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
  "gobot/input"
  "gobot/indexer"
  "gobot/indexmap"
)


func main() {
  indexMap := indexmap.New();
  list := input.Input("input/urls")
  for e := list.Front(); e != nil; e = e.Next() {
    url := e.Value.(string)
    list := indexer.Indexing(url)
    indexMap.AddIndexStringList(list)
  }
  //reducer.Print()
}
