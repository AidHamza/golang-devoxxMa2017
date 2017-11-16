package main

import "fmt"

func main() {

    // synchronously call of a function
    loopPrint("called this directly", 3)

    // call using go routine
    go loopPrint("<---- calling this from goroutine ---->", 10)

    // call using Go Routine
    go func(msg string) {
        fmt.Println(msg)
    }("<---- Another yet Simple Routine ---->")

    var input string
    fmt.Scanln(&input)
    fmt.Println("program exited")
}

func loopPrint(from string, times int) {
    for i := 0; i < times; i++ {
        fmt.Println(from, ":", i)
    }
}