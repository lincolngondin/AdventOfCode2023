package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "fmt"
)

func GetHistoryFromFile(fileName string) [][]int {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var data [][]int = make([][]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
        valuesStr := strings.Split(line, " ")
        var value int
        var temp []int = make([]int, 0)
        for _, v := range(valuesStr) {
             fmt.Sscanf(v, "%d", &value)
             temp = append(temp, value)
        }
        data = append(data, temp)
    }
    return data
}
