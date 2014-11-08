package main

import (
    "fmt"
    "gobot/reducer"
    "gobot/parser"
)

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
    worker := Worker {
        ID:             id,
        Work:           make(chan WorkRequest),
        WorkerQueue:    workerQueue,
        QuitChan:       make(chan bool)}

    return worker
}

type Worker struct {
    ID          int
    Work        chan WorkRequest
    WorkerQueue chan chan WorkRequest
    QuitChan    chan bool
}

func (w Worker) Start() {
    go func() {
        for {
            WorkerQueue <-w.Work

            select {
            case work := <-w.Work:
                result, err := parser.Indexing(work.url)
                reducer.Reduce(result)
            case <-w.QuitChan:
                fmt.Printf("worker %d stopping\n", w.ID)
                return
            }
        }
    }()
}

func (w Worker) Stop() {
    go func() {
        w.QuitChan <- true
    }()
}


