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
  urls string
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
  w.Work <- WorkRequest {urls: "", term: true}
}

func NewWorker(id int) Worker {
  worker := Worker {
    ID: id,
    Work: make(chan WorkRequest),
  }
  return worker
}

func GobotPmap() {
  worker := NewWorker(0)
  worker.Start()

  list := input.Input("input/urls")
  for e := list.Front(); e != nil; e = e.Next() {
    url := e.Value.(string)

    worker.Work <- WorkRequest {urls: url, term: false}
  }
  worker.Stop();
}
