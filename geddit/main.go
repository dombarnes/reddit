package main

import (
    "fmt"
    "github.com/dombarnes/reddit"
    "log"
)

func main() {
    items, err := reddit.Get("deadbedrooms")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
