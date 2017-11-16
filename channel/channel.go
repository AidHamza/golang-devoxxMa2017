package main

import "fmt"

func myFunc(done chan string) {
    // Doing something in parallel
    for i := 0; i < 10; i++ {
      fmt.Println(i)
    }
    fmt.Println("Hey! I do useless stuff!")
    done <- "I'm done!" // We send a message on the channel
}

func main() {
    done := make(chan string)

    go myFunc(done)

    msg := <- done // Block until we receive a message on the channel
    fmt.Println(msg)

    fmt.Println("Message received, you were indeed useless..")
}