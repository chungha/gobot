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
    StartDispatcher(1)

    for {
        fmt.Print("Let me know the url :")
        Collector(os.Stdin)
        fmt.Print("\n")
        reducer.Print()
        fmt.Print("\nCommand :")
        Command()
        fmt.Println()
    }
}
