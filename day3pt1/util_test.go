package main

import (
    "testing"
    "bytes"
)

func TestGetValueInPosition(t *testing.T){
    var str string = "467..114.3"
    var bf *bytes.Buffer = bytes.NewBufferString(str)
    var test []byte = bf.Bytes()
    testSuite := []int{467, 467, 467, 0, 0, 114, 114, 114, 0, 3}
    for i, v := range(testSuite) {
        result := getValueInPosition(test, i)
        if result != v {
            t.Error("Error!!")
        }
    }
}
