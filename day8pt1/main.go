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


func takeSteps(mapa map[string]MapNode, inst string) int {
    steps := 0
    instSize := len(inst)
    instPos := 0

    var actualPos string = "AAA"
    for steps = 0; actualPos != "ZZZ"; steps++ {
        dir := rune(inst[instPos])
        if dir == 'L' {
            actualPos = mapa[actualPos].left
        } else {
            actualPos = mapa[actualPos].right
        }

        instPos++
        if instPos >= instSize {
            instPos = 0
        }
    }
    return steps
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
    file, err := os.Open("input2.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    scan := bufio.NewScanner(file)

    mapa, instructions := getMapAndInstructions(scan)
    steps := takeSteps(mapa, instructions)
    fmt.Printf("It is required %d steps!\n", steps)


}
