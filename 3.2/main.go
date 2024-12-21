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

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/3.2/input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    LineStartEnabled := regexp.MustCompile(`(?U)^.*don\'t\(\)`)
    OtherEnabled     := regexp.MustCompile(`(?U)do\(\).*(don't\(\)|$)`)
    enabledMatches := make([]string,0)
    loopCounter := 0
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println(line)
        matchStartLine := LineStartEnabled.FindAllString(line, -1)
        matchOther     := OtherEnabled.FindAllString(line, -1)
        fmt.Println("---------------LINE START---------------")
        fmt.Println(matchStartLine)
        fmt.Println("---------------OTHER MATCH---------------")
        fmt.Println(matchOther)
        if loopCounter == 0 {
            enabledMatches = append(enabledMatches, matchStartLine...)
        }
        enabledMatches = append(enabledMatches, matchOther...)
        fmt.Println("---------------ENABLED MATCHES SLICE INNER LOOP---------------")
        fmt.Println(enabledMatches)
        loopCounter++
    }

    fmt.Println("---------------ENABLED MATCHES SLICE---------------")
    fmt.Println(enabledMatches)


    rg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
    allMatches := make([]string,0)
    for i := range enabledMatches {
        fmt.Println("---------------SINGLE ENABLED MATCH---------------")
        fmt.Println(enabledMatches[i])
        match := rg.FindAllString(enabledMatches[i], -1)
        fmt.Println("---------------MATCH SLICE---------------")
        fmt.Println(match)
        allMatches = append(allMatches, match...)
    }

    fmt.Println(allMatches)

    numsToMultiply := regexp.MustCompile(`[0-9]+`)
    var result int
    for i := range allMatches {

        strNums := numsToMultiply.FindAllString(allMatches[i], -1)
        //fmt.Println(strNums)

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

}
