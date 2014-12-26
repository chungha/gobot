/**
 * The main function of the gobot.
 * Input : A text file containing urls.
 * Output : The time of indexing.
 */
package main

import (
  "fmt"
  "gobot/input"
)

type WorkRequest struct {
  urls []string
  term bool
}

type Worker struct {
  ID int
  Work chan WorkRequest
}

func (w Worker) Start() {
  go func() {
    for {
      work := <-w.Work

      if work.term {
        fmt.Printf("worker %d stopping\n", w.ID)
        return
      }

      fmt.Printf("input - %s\n", work.urls)
    }
  }()
}

func (w Worker) Stop() {
  w.Work <- WorkRequest {urls: nil, term: true}
}

func NewWorker(id int) Worker {
  worker := Worker {
    ID: id,
    Work: make(chan WorkRequest),
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
    workers[ths].Work <- WorkRequest {urls: url_array[index:index+offset], term: false}
    index = index + offset
  }
  workers[num_thread-1].Work <- WorkRequest {urls: url_array[index:url_len], term: false}

  for i:=0; i<num_thread; i++ {
    workers[i].Stop()
  }
}
