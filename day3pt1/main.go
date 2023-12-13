package main

import (
    "os"
)

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()
}
