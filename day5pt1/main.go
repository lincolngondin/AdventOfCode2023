package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "fmt"
)

type Range struct {
    destinationRangeStart int
    sourceRangeStart int
    rangeLength int
}

type Input struct {
    seeds []int
    maps map[string][]Range
}

func extractRange(line string) Range {
    rg := Range{0, 0, 0}
    fmt.Sscanf(line, "%d %d %d", &rg.destinationRangeStart, &rg.sourceRangeStart, &rg.rangeLength)
    return rg
}

func extractSeeds(line string) []int {
    temp := strings.Split(line, ": ")
    numbers := make([]int, 0, 10)
    numbersString := strings.Split(temp[1], " ")
    for _, numberStr := range(numbersString){
        value := 0
        fmt.Sscanf(numberStr, "%d", &value)
        numbers = append(numbers, value)
    }
    return numbers
}

func extractMapName(line string) string {
    var mp string
    fmt.Sscanf(line, "%s map:", &mp)
    return mp
}

func convertSource(mappedValue []Range, sourceValue int) int {

}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)
    input := Input{[]int{}, map[string][]Range{}}
    var insideMap bool = false
    var mapName string

    for fileScanner.Scan(){
        line := fileScanner.Text()
        if strings.Contains(line, "seeds"){
            input.seeds = extractSeeds(line)
        } else if strings.Contains(line, "map"){
            mapName = extractMapName(line)
            insideMap = true;
        } else if line == "\n" {
            insideMap = false;
        } else {
            if insideMap && len(line) > 5 {
                input.maps[mapName] = append(input.maps[mapName], extractRange(line))
            }
        }
    }
    order := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
    for mpn := range(order) {
        for i := 0; i < len(input.seeds); i++ {
            for v := range(input.maps[mpn]){
                if (input.seeds[i] >= v.sourceRangeStart) && (v <= v.sourceRangeStart+v.rangeLength) {
                    input.seeds[i] = v.destinationRangeStart + (input.seeds[i] - v.sourceRangeStart)
                    break
                }
            }
        }
    }


}
