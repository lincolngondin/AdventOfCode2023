package main

import (
	"log"
	"os"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    var (
        file *os.File
        err error
    )
    file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file!")
	}
	defer file.Close()

    scan := bufio.NewScanner(file)

    var numbers map[string]int = map[string]int{"one":1, "two":2, "three":3, "four":4, "five":5, "six":6, "seven":7, "eight":8, "nine":9}
    var sum int = 0
    
    for scan.Scan() {
        var line string = scan.Text()
        var firstDigit, lastDigit int = -1, -1
        var digits []int = make([]int, len(line))
        for i := 0; i < len(line); i++{
            digits[i] = -1
        }

        // add digits text
        for key, value := range(numbers) {
            count := strings.Count(line, key)
            if count == 1 {
                idx := strings.Index(line, key)
                digits[idx] = value
            } else if count >= 2 {
                idx := strings.Index(line, key)
                idx2 := strings.LastIndex(line, key)
                digits[idx] = value
                digits[idx2] = value
            } 
        }

        // add digits numbers
        read := strings.NewReader(line)
        idx := 0
        for{
            digitByte, err := read.ReadByte()
            if err != nil {
                break
            }
            if digitByte <= '9' && digitByte >= '0' {
                digits[idx] = int(digitByte) - 48
            }
            idx++
        }

        for _, v := range(digits) {
            if v != -1 {
                if firstDigit == -1 {
                    firstDigit = v
                }
                lastDigit = v
            }
        }

        sum += firstDigit*10 + lastDigit

    }
    fmt.Printf("The sum is %d!\n", sum)


}
