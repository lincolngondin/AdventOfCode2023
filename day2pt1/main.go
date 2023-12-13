package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
    "strings"
)

type Game struct {
    ID int
    sets []map[string]int
}

func getGameFromLine(s string) Game {
    
    var ID int = 0
    set := []map[string]int{}
    Strs := strings.Split(s, ":")
    fmt.Sscanf(Strs[0], "Game %d", &ID)
    sets := strings.Split(Strs[1], ";")
    for _, tset := range(sets){
        set = append(set, getGameSet(tset))
    }
    return Game{ID, set} 
    
}

func getGameSet(s string) map[string]int {
    gset := map[string]int{
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    cubes := strings.Split(s, ",")
    for _, cube := range(cubes) {
        q := 0
        var color string = ""  
        fmt.Sscanf(cube, "%d %s", &q, &color)
        gset[color] = q
    
    }
    return gset

}

func isPossible(game *Game) bool {
    qRed, qGreen, qBlue := 12, 13, 14
    for _, set := range(game.sets){
        if set["blue"] > qBlue || set["red"] > qRed || set["green"] > qGreen {
            return false
        }
    }
    return true
}

func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file!")
    }
    defer file.Close()

    var lineScanner *bufio.Scanner = bufio.NewScanner(file)
    var line string
    var sum int = 0
    for lineScanner.Scan() {
        line = lineScanner.Text()
        gm := getGameFromLine(line)
        if isPossible(&gm) {
            sum += gm.ID
        }
    }
    fmt.Printf("The sum is %d!\n", sum)
    

}
