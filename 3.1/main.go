package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "regexp"
    "strconv"
)

func main () {

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/3.1/input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    rg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
    allMatches := make([]string,0)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
        match := rg.FindAllString(line, -1)
        allMatches = append(allMatches, match...)
    }

    fmt.Println(allMatches)

    numsToMultiply := regexp.MustCompile(`[0-9]+`)
    var result int
    for i := range allMatches {

        strNums := numsToMultiply.FindAllString(allMatches[i], -1)
        fmt.Println(strNums)

        nums := make([]int, 0)
        for j := range strNums {
            n, err := strconv.Atoi(strNums[j])
            if err != nil {
                log.Fatalf("unable to convert")
            }
            nums = append(nums, n)
        }
        fmt.Println(nums)

        multRes := nums[0] * nums[1]
        fmt.Println(multRes)
        result = result + multRes
    }

    fmt.Println(result)
    //379055652 to HIGH

}
