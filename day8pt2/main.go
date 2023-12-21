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
    var endPoints []string = make([]string, 0)
    for k, _ := range(mapa) {
        if isStartingPoint(k){
            startingPoints = append(startingPoints, strings.Clone(k))
        }
        if isFinalNode(k){
            endPoints = append(endPoints, k)
        }
    }
    pointsSize := len(startingPoints)

    
    for steps = 0; !nodesEndsWithZ(startingPoints); steps++{
        dir := rune(inst[instPos])
        for i := 0; i < pointsSize; i++ {
            if dir == 'L' {
                startingPoints[i] = mapa[startingPoints[i]].left
            } else {
                startingPoints[i] = mapa[startingPoints[i]].right
            }
        }

        
        instPos++
        if instPos >= instSize {
            instPos = 0
        }
    }
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
        mapa[parts[0]] = MapNode{left, right}
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
