/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
  "fmt"
  "gobot/indexer"
  "gobot/reducer"
)


func main() {
  list := indexer.Indexing("http://daum.net")
  Reduce(list)
}
