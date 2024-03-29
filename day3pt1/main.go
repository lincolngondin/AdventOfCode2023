package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

/*
   2              
[467..112..]
   i

input array of string, index base (string[], index)
    if array[index] is not an number
        return 0
    while array[index] is an number
        index++
        if index >= len(array)
            break

    index--
    mult = 1
    value = 0
    while array[index] is an number and index >= 0
        value += to_number(array[index]) * mult
        mult *= 10
        index--
    return value
*/

/*
input array of array of bytes (bytes[][])
    for any value if it is a symbol sum the surroundings
*/

func is_number(v byte) bool {
    return v <= '9' && v >= '0'
}

func to_number(v byte) int {
    if is_number(v) {
        return int(v) - 48
    }
    log.Fatal("Value is not a number!")
    return 0
}

func sumAllParts(data [][]byte) int{
    sum := 0
    for y := 1; y < 139; y+=1 {
        for x := 1; x < 139; x++ {
             if !is_number(data[y][x]) && data[y][x] != '.'{
                sum+= getValueInPosition(data[y], x-1)
                sum+= getValueInPosition(data[y], x+1)
                if is_number(data[y+1][x]) {
                    sum+= getValueInPosition(data[y+1], x)
                } else{
                    sum+= getValueInPosition(data[y+1], x-1)
                    sum+= getValueInPosition(data[y+1], x+1)
                }
                

                if is_number(data[y-1][x]) {
                    sum+= getValueInPosition(data[y-1], x)
                } else {
                    sum+= getValueInPosition(data[y-1], x-1)
                    sum+= getValueInPosition(data[y-1], x+1)
                }
            }
        }
    }
    return sum
}

func getValueInPosition(line []byte, index int) int {
    if !is_number(line[index]) {
        return 0
    }
    for is_number(line[index]) {
        index++
        if index >= len(line){
            break
        }
    } 
    index--
    var value int = 0
    var mult int = 1
    for is_number(line[index]) {
        value += to_number(line[index]) * mult 
        mult *= 10
        index--
        if index < 0 {
            break
        }
    }
    return value
}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    scan := bufio.NewScanner(file)

    var data [][]byte = make([][]byte, 140)
    for idx:= 0; scan.Scan(); idx++ {
        line := scan.Bytes()
        data[idx] = make([]byte, 140)
        for i := 0; i < len(line); i++ {
            data[idx][i] = line[i]
        }
    }
    sum := sumAllParts(data)
    fmt.Println("The sum of all parts numbers is ", sum)

}
