package main

import (
    "os"
    "fmt"
    "gobot/reducer"
    "bufio"
)

func Command() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        cmd := scanner.Text()
        if cmd == "." {
            break
        }
        fmt.Println(reducer.Query(cmd))
    }
}

func main() {
    fmt.Println("Starting the dispatcher")
    StartDispatcher(4)

    for {
        Collector(os.Stdin)
        Command()
    }
}
