package main

import (
    "testing"
)

func TestGetMatchingNumberQuantity(t *testing.T){
    testSuite := []Card{
        {1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}},
        {2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}},
    }
    expectedSuite := []int{4, 2}
    for i, test := range(testSuite) {
        result := getMatchingNumberQuantity(test)
        if result != expectedSuite[i] {
            t.Error("Erro")
        }

    }
}

func TestProcessScratchCards(t *testing.T) {
    testSuite := []CardInstance{
        {Card{1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}, 0},
        {Card{2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}}, 0},
        {Card{3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}, 0},
        {Card{4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}}, 0},
        {Card{5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}}, 0},
        {Card{6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}, 0},
    }
    processScratchCards(testSuite)
    result := getInstancesQuantity(testSuite)
    if result != 30 {
        t.Error("Get: ", result, " Expect: ", 30, " Erro!")
    }
}
