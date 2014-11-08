package main

import (
    "fmt"
    "bufio"
    "io"
)

var WorkQueue = make(chan WorkRequest, 100)

func request(url string) {
    work := WorkRequest{url: url}
    WorkQueue <- work
    fmt.Println("requested : ", url)
}

func Collector(reader io.Reader) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        url := scanner.Text()
        if url == "." {
            break;
        }
        request(url)
    }
}
