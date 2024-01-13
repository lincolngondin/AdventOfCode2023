package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
    "strings"
)

type MapNode struct {
    left string
    right string
}

func nodesEndsWithZ(nodes []string) bool {
    for _, v := range(nodes) {
        if !isFinalNode(v){
            return false
        }
    }
    return true
}
func CountNodesEndsWithZ(nodes []string) int {
    count := 0
    for _, v := range(nodes) {
        if isFinalNode(v){
            count++
        }
    }
    return count
}

func takeSteps(mapa map[string]MapNode, inst string) int {
    steps := 0
    instSize := len(inst)
    instPos := 0

    var startingPoints []string = make([]string, 0)
    var values []int = make([]int, 0)
    for k, _ := range(mapa) {
        if isStartingPoint(k){
            startingPoints = append(startingPoints, strings.Clone(k))
            values = append(values, 0)
        }
    }
    pointsSize := len(startingPoints)

    var step int = 0
    fmt.Println(values)

    for i := 0; i < pointsSize; i++ {
        
        instPos = 0
        step = 0
        for {
            dir := rune(inst[instPos])
            if dir == 'L' {
                startingPoints[i] = mapa[startingPoints[i]].left
            } else {
                startingPoints[i] = mapa[startingPoints[i]].right
            }
            step++
            if isFinalNode(startingPoints[i]) {
                break
            }
            instPos++
            if instPos >= instSize {
                instPos = 0
            }
        }
        values[i] = step
    }
    fmt.Println(startingPoints)
    fmt.Println(values)
    return steps
}

func isStartingPoint(node string) bool {
    return node[2] == 'A'
}
func isFinalNode(node string) bool {
    return node[2] == 'Z'
}

func getMapAndInstructions(scan *bufio.Scanner) (map[string]MapNode, string) {
    scan.Scan()
    var instructions string = scan.Text()
    scan.Scan()

    var mapa map[string]MapNode = make(map[string]MapNode)

    for scan.Scan(){
        line := scan.Text()
        parts := strings.Split(line, " = ")
        parts2 := strings.Split(parts[1], ",")
        var left, right string
        fmt.Sscanf(parts2[0], "(%s", &left)
        fmt.Sscanf(strings.ReplaceAll(parts2[1], ")", ""), " %s)", &right)
        mapa[strings.Clone(parts[0])] = MapNode{strings.Clone(left), strings.Clone(right)}
    }
    return mapa, instructions
}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    scan := bufio.NewScanner(file)

    mapa, instructions := getMapAndInstructions(scan)
    steps := takeSteps(mapa, instructions)
    fmt.Printf("It is required %d steps!\n", steps)

}
