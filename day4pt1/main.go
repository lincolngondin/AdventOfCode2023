package main

import (
    "os"
    "bufio"
    "strings"
    "fmt"
    "log"
    "slices"
)

type Card struct {
    cardNumber int
    winningNumbers []int
    numbers []int
}

func getCard(line string) Card {
    number := 0
    var winningNumbers []int = make([]int, 10)
    var numbers []int = make([]int, 25)
    var temp []string = strings.Split(line, ":")
    fmt.Sscanf(temp[0], "Card %d:", &number)
    var strs []string = strings.Split(temp[1], "|")
    var wn []string = strings.Split(strs[0], " ")
    var mn []string = strings.Split(strs[1], " ")
    var idx int = 0
    for _, v := range(wn) {
        value := 0
        fmt.Sscanf(v, "%d", &value)
        if value != 0{
            winningNumbers[idx] = value
            idx++;
        }
    }
    idx = 0
    for _, v := range(mn) {
        value := 0
        fmt.Sscanf(v, "%d", &value)
        if value != 0{
            numbers[idx] = value
            idx++;
        }
    }

    return Card{number, winningNumbers, numbers} 
}

func getWorthPoints(card Card) int {
    var point int = 0
    for _, mynumber := range(card.numbers) {
        if slices.Contains(card.winningNumbers, mynumber) {
            if point == 0 {
                point = 1
            } else{
                point *= 2
            }
        }
    }
    return point
}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    var scannedFile *bufio.Scanner = bufio.NewScanner(file)
    var line string
    var sumPoints int = 0
    for scannedFile.Scan() {
         line = scannedFile.Text()
         var cd Card = getCard(line)
         sumPoints+=getWorthPoints(cd)
    }
    fmt.Printf("There are %d points in total!\n", sumPoints)
} 
