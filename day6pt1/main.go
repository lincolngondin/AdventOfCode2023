package main

import (
    "fmt"
)

/*
calculate_distance(hold_time, max_time)
    return (max_time-hold_time) * hold_time
input time and record distance (time, record) output ways_to_win int
    ways = 0
    for i = 0 until time
        if calculate_distance(i, time) > record
            ways++ 
        
    return ways

*/

func calculate_distance(hold_time int, max_time int) int {
    return (max_time - hold_time) * hold_time
}

func race_ways_to_win(time int, record int) int {
    ways := 0
    for i := 0; i <= time; i++ {
        if calculate_distance(i, time) > record {
            ways++
        }
    }
    return ways
}

func main(){
    var times []int = []int{56, 71, 79, 99}
    var distances []int = []int{334, 1135, 1350, 2430}
    var sum int = 1
    for idx := 0; idx < len(times); idx++ {
        sum *= race_ways_to_win(times[idx], distances[idx])
    }
    fmt.Println("The answer is ", sum)
}
