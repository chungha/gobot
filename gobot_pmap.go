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

type WorkRequest struct {
  urls []string
}

type Worker struct {
  ID int
  Work chan WorkRequest
  Done chan bool
}

func (w Worker) Start() {
  go func() {
    for {
      work := <-w.Work

      fmt.Printf("ID:%d = input(%d) - %s\n", w.ID, len(work.urls), work.urls)
      var indexMap = indexer.IndexingUrlsToMap(work.urls)
      merger.Merge(&globalIndexMap, indexMap)

      w.Done <- true
    }
  }()
}

func NewWorker(id int) Worker {
  worker := Worker {
    ID: id,
    Work: make(chan WorkRequest),
    Done: make(chan bool),
  }
  return worker
}

func GobotPmap(num_thread int) {
  if num_thread <= 0 {
    fmt.Printf("Thread should be more than 1")
    return
  }

  workers := make([]Worker, num_thread)
  for i:=0; i<num_thread; i++ {
    worker := NewWorker(i)
    worker.Start()
    workers[i] = worker
  }

  list := input.Input("input/urls")
  var url_array []string
  for e := list.Front(); e != nil; e = e.Next() {
    url := e.Value.(string)
    url_array = append(url_array, url)
  }

  url_len := len(url_array)
  offset := url_len/num_thread
  index := 0
  for ths:=0; ths<(num_thread-1); ths++ {
    workers[ths].Work <- WorkRequest {urls: url_array[index:index+offset]}
    index = index + offset
  }
  workers[num_thread-1].Work <- WorkRequest {urls: url_array[index:url_len]}

  for i:=0; i<num_thread; i++ {
    <-workers[i].Done
  }
}
