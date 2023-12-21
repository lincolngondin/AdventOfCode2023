package main

import (
    "testing"
    "strings"
)

func TestGetType(t *testing.T){
    inputs := []string{"32T3K", "T55J5", "KK677", "KTJJT", "QQQJA"}
    expect := []string{"One pair", "Three of a kind", "Two pair", "Two pair", "Three of a kind"}
    for i, v := range(inputs){
        result := getType(v)
        if strings.Compare(expect[i], result) != 0 {
            t.Error("For: ", v, " Get: ", result, " Expected: ", expect[i])
        }
    }
}

func TestHandCompare(t *testing.T){
    result := handCompare("T55J5", "QQQJA")
    if result != -1 {
        t.Error("Error!")
    }
    result2 := handCompare("KK677", "KTJJT")
    if result2 != 1 {
        t.Error("Error!")
    }
}

func TestProgram(t *testing.T){
    inputTest := []inputType{
        {"32T3K", "One pair", "", 765},
        {"T55J5", "Three of a kind", "", 684},
        {"KK677", "Two pair", "", 28},
        {"KTJJT", "Two pair", "", 220},
        {"QQQJA", "Three of a kind","", 483},
    }
    for i := 0; i < len(inputTest); i++ {
        inputTest[i].bestHandType = getBetterHandType(inputTest[i].hand)
    }

    orderHands(inputTest)
    sum := sumAllWins(inputTest)
    if sum != 5905 {
        t.Error(" Get: ", sum, " Expected: ", 5905)
    }

}
