/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
	"gobot/backend_indexer/indexer"
	"gobot/backend_indexer/indexmap"
	"gobot/backend_indexer/input"
)

func GobotSequential() {
	indexMap := indexmap.NewSync()
	list := input.Input("input/urls")
	for e := list.Front(); e != nil; e = e.Next() {
		url := e.Value.(string)
		list := indexer.Indexing(url)
		indexMap.AddIndexStringList(list)
	}
	//reducer.Print()
}
