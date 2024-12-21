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

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/1.2/input.txt")
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

    file, err = os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/1.2/input2.txt")
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

    var result int = 0
    var counter int
    for _, left := range(IDsLeft) {
        counter = 0
        for _, right := range (IDsRight) {
            if left == right {
                counter = counter + 1
            }
        }
        result = result + (left * counter)
    }

    fmt.Println(result)
}
