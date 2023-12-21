package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
    "strings"
)

type inputType struct {
    hand string
    handType string
    bid int
}

func getType(hand string) string {
    var values map[rune]int = make(map[rune]int,5)
    for _, v := range(hand){
        values[v]++
    }
    switch len(values) {
        case 1:
            return "Five of a kind"
        case 2:
            for _,v := range(values){
                if v == 4 {
                    return "Four of a kind"
                }             
            }
            return "Full house"
        case 3:
            for _,v := range(values){
                if v == 3 {
                    return "Three of a kind"
                } 
            }
            return "Two pair"
        case 4:
            return "One pair"
        case 5:
            return "High card"
    }
    return ""
}

/*
input two strings to compare return 0 if equal -1 if less 1 is bigger
    for each carachter c1, c2 in each string 
        if c1 != c2
            if c1 > c2
                return 1
            else 
                return -1
    return 0
*/
var cardStrength map[rune]int = map[rune]int {
    'A': 13,
    'K': 12,
    'Q': 11,
    'J': 10,
    'T': 9,
    '9': 8,
    '8': 7,
    '7': 6,
    '6': 5,
    '5': 4,
    '4': 3,
    '3': 2,
    '2': 1,
}

func handCompare(s1, s2 string) int {
    for i, c1 := range(s1) {
        if c1 != rune(s2[i]){
            if cardStrength[c1] > cardStrength[rune(s2[i])] {
                return 1
            } else {
                return -1
            }
        }
    }
    return 0
}


func orderHands(hands []inputType){
    var n int = len(hands)
    var handTypes map[string]int = map[string]int{
        "Five of a kind": 7,
        "Four of a kind": 6,
        "Full house": 5,
        "Three of a kind": 4,
        "Two pair": 3,
        "One pair": 2,
        "High card": 1,
    }
    for i := n-1; i > 0; i-- {
        for b := 0; b < i; b++ {
            //se forem do mesmo tipo
            if strings.Compare(hands[b].handType, hands[b+1].handType) == 0{
                if handCompare(hands[b].hand, hands[b+1].hand) > 0 {
                    temp := hands[b]
                    hands[b] = hands[b+1]
                    hands[b+1] = temp
                }
            } else {
                if handTypes[hands[b].handType] > handTypes[hands[b+1].handType] {
                    temp := hands[b]
                    hands[b] = hands[b+1]
                    hands[b+1] = temp
                }
            }
        }
    }
}

func sumAllWins(data []inputType) int {
    sum := 0
    for i, v := range(data){
        wins := (i+1) * v.bid
        sum+=wins
    }
    return sum
}

func main(){

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file")
    }
    defer file.Close()
    scan := bufio.NewScanner(file)
    var input []inputType = make([]inputType, 0, 1000)
    for scan.Scan(){
        line := scan.Text()
        it := inputType{"", "", 0}
        fmt.Sscanf(line, "%s %d", &it.hand, &it.bid)
        it.handType = getType(it.hand)
        input = append(input, it)
    }

    orderHands(input)
    sum := sumAllWins(input)

    fmt.Println("The total winnings are ", sum)
    
}
