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

type CardInstance struct {
    card Card
    instances int
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

func getMatchingNumberQuantity(card Card) int {
    var quantity int = 0
    for _, mynumber := range(card.numbers) {
        if slices.Contains(card.winningNumbers, mynumber) {
            quantity++
        }
    }
    return quantity
}

func getInstancesQuantity(cards []CardInstance) int {
    var q int = 0
    for _, card := range(cards) {
        q += card.instances
    }
    return q
}

func processFromIndex(cards []CardInstance, idx int, limit int){
    matches := getMatchingNumberQuantity(cards[idx].card)
    for i := idx+1; i < limit && i <= idx+matches; i++ {
        cards[i].instances++
        processFromIndex(cards, i, limit)
    }
}

func processScratchCards(cards []CardInstance) {
    var limit int = len(cards)
    for idx := 0; idx < limit; idx++{
        cards[idx].instances++
        processFromIndex(cards, idx, limit)
    }
}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    var line string
    var cards []CardInstance = make([]CardInstance, 0, 218)

    var scannedFile *bufio.Scanner = bufio.NewScanner(file)
    for scannedFile.Scan() {
         line = scannedFile.Text()
         var cd Card = getCard(line)
         cards = append(cards, CardInstance{cd, 0})
    }
    processScratchCards(cards)
    fmt.Printf("There are %d scratchcards in the end!\n", getInstancesQuantity(cards))
} 
