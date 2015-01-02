/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
  "fmt"
  "gobot/input"
  "gobot/indexmap"
  "gobot/indexer"
  "gobot/merger"
)

var globalIndexMap = indexmap.NewSync()

type Worker struct {
  ID int
  Work chan string
  Done chan bool
}

func (w Worker) Start() {
  go func() {
    var indexMap = indexmap.New()
    for {
      url := <-w.Work

      if len(url) > 0 {
        indexer.IndexingToMap(url, indexMap)
      } else {
        merger.Merge(&globalIndexMap, indexMap)
        w.Done <- true
        break;
      }
    }
  }()
}

func NewWorker(id int, workChannel chan string) Worker {
  worker := Worker {
    ID: id,
    Work: workChannel,
    Done: make(chan bool),
  }
  return worker
}

func GobotPmap(num_thread int) {
  if num_thread <= 0 {
    fmt.Printf("Thread should be more than 1")
    return
  }

  var workChannel = make(chan string)
  workers := make([]Worker, num_thread)
  for i:=0; i<num_thread; i++ {
    worker := NewWorker(i, workChannel)
    worker.Start()
    workers[i] = worker
  }

  list := input.Input("input/urls")
  for e := list.Front(); e != nil; e = e.Next() {
    url := e.Value.(string)
    workChannel <- url
  }
  for i := 0; i < num_thread; i++ {
    workChannel <- ""
  }

  for i := 0; i < num_thread; i++ {
    <-workers[i].Done
  }
}
