package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "sort"
)

func main() {

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/1/input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    var IDsLeft []int
    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }
    for scanner.Scan() {
        ID, err := strconv.Atoi(scanner.Text())
        if err != nil {
           log.Fatalf("unable conver to int", err)
        }
        IDsLeft = append(IDsLeft, ID)
    }

    file, err = os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/1/input2.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    var IDsRight []int
    scanner = bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }
    for scanner.Scan() {
        ID, err := strconv.Atoi(scanner.Text())
        if err != nil {
           log.Fatalf("unable conver to int", err)
        }
        IDsRight = append(IDsRight, ID)
    }


    sort.Sort(sort.IntSlice(IDsLeft))
    sort.Sort(sort.IntSlice(IDsRight))

    fmt.Println("--------------------------------------------------------------------")
    fmt.Println(IDsRight)
    fmt.Println(IDsLeft)

    var plaseHolder int
    var result int = 0
    for i := range(IDsLeft) {
        if IDsLeft[i] > IDsRight[i] {
            plaseHolder = IDsLeft[i] - IDsRight[i]
        } else {
            plaseHolder = IDsRight[i] - IDsLeft[i]
        }
        result = result + plaseHolder
    }

    fmt.Println(result)
}
