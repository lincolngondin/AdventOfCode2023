package main

import (
    "fmt"
)


func allZeroes(sequence []int) bool {
    for _, v := range(sequence){
        if v != 0 {
            return false
        }
    }
    return true
}

func getExtrapoletedValue(data []int) int {
    values := [][]int{data}
    for b:=0; !allZeroes(values[b]); b++{
        temp := make([]int, 0)
        for i := 0; i < len(values[b])-1; i++ {
            temp = append(temp, values[b][i+1]-values[b][i])
        }
        values = append(values,temp)
    }
    historyValues := 0
    for i := len(values)-1; i > 0; i-- {
        historyValues = (values[i-1][0]) - historyValues
    }

    return historyValues
}

func main(){
    data := GetHistoryFromFile("input.txt")
    sum := 0
    for _, v := range(data){
        sum += getExtrapoletedValue(v)
    }
    fmt.Println("The sum of the extrapolated values is ", sum)
}
