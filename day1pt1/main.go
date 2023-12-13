package main

import (
	"fmt"
	"log"
	"os"
    "bufio"
    "strings"
    "io"
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
    var sum int = 0

    for scan.Scan() {
        read := strings.NewReader(scan.Text())
        var dig1, dig2 int
        var q int = 0
        for {
            c, err := read.ReadByte()
            if err == io.EOF {
                break
            }
            if c <= '9' && c >= '0' {
                if q == 0{
                    dig1 = int(c) - 48
                    dig2 = int(c) - 48
                    q++;
                } else{
                    dig2 = int(c) - 48
                }
            }
        }
        sum += dig1*10 + dig2
    }
    fmt.Printf("The sum is %d!\n", sum)



}
