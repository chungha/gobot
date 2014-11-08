package main

import (
    "os"
    "fmt"
    "bufio"
    "container/list"
)

type WorkRequest struct {
    url string
}

var WorkQueue = make(chan WorkRequest, 100)

func ReadURLs(filename string) list.List {
    var urls list.List

    f, err := os.Open(filename)
    if err != nil {
        return urls
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        url := scanner.Text()
        urls.PushBack(url)
    }

    return urls
}

func Request(url string) {
    work := WorkRequest{url: url}
    WorkQueue <- work
    fmt.Println("requested : ", url)
}

func RequestAll(urls list.List) {
    for url := urls.Front(); url != nil; url = url.Next() {
        if v, ok := url.Value.(string); ok {
            Request(v)
        }
    }
}

func main() {
    if len(os.Args) == 2 {
        input_file := os.Args[1]
        var urls = ReadURLs(input_file)
        RequestAll(urls)
    }

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        url := scanner.Text()
        if url == "." {
            break;
        }
        Request(url)
    }
}
