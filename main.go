package main

import "fmt"
import "gobot/parser"

func main() {
    html, err := parser.Download("http://golang.org/")
	if (err != nil) {
	  fmt.Printf("[ERROR] %s", err)
    } else {	
	  fmt.Printf(html)
	}
	fmt.Printf("hello, world\n")
}