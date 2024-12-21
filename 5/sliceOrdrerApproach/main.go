package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "slices"
)

func main () {

    file, err := os.Open("./input.order")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    ruleSlice := make([]int,0)
    nums := make([]int,2)
    var e error

    for scanner.Scan() {
        line := scanner.Text()
        strs := strings.Split(line, "|")
        for i := range nums {
            nums[i], e = strconv.Atoi(strs[i])
            if e != nil {
                panic(e)
            }
        }
        //fmt.Println(nums)
        if !(slices.Contains(ruleSlice, nums[0])) && !(slices.Contains(ruleSlice, nums[1])) {
            ruleSlice = append(ruleSlice, nums[0])
            ruleSlice = append(ruleSlice, nums[1])
            fmt.Println(ruleSlice)
            fmt.Println("--------")
            //fmt.Println("CONTINUE")
            continue
        }
        if (slices.Contains(ruleSlice, nums[0])) && !(slices.Contains(ruleSlice, nums[1])) {
            ruleSlice = append(ruleSlice, nums[1])
            fmt.Println(ruleSlice)
            fmt.Println("--------")
            //fmt.Println("CONTINUE")
            continue
        }
        if !(slices.Contains(ruleSlice, nums[0])) && (slices.Contains(ruleSlice, nums[1])) {
            //fmt.Println("HERE")
            for i := range ruleSlice {
                if ruleSlice[i] == nums[1] {
                    tempSlice := make([]int, len(ruleSlice[:i]))
                    tempSlice2 := make([]int, len(ruleSlice[i:]))
                    copy(tempSlice, ruleSlice[:i])
                    copy(tempSlice2, ruleSlice[i:])
                    //fmt.Println(tempSlice)
                    //fmt.Println(tempSlice2)
                    ruleSlice = append(tempSlice, nums[0])
                    ruleSlice = append(ruleSlice, tempSlice2[0:]...)
                    break
                }
                continue
            }
            fmt.Println(ruleSlice)
            fmt.Println("--------")
            continue
        }

        if (slices.Contains(ruleSlice, nums[0])) && (slices.Contains(ruleSlice, nums[1])) {
            firstIDX := -1
            secondIDX := -1
            for i := range ruleSlice {
                if ruleSlice[i] == nums[0] {
                    firstIDX = i
                }
                if ruleSlice[i] == nums[1] {
                    secondIDX = i
                }
                if (firstIDX != -1) && (secondIDX != -1) {
                    if firstIDX < secondIDX {
                        //fmt.Println("BREAK")
                        break
                    }
                    tempSlice := make([]int, len(ruleSlice[:secondIDX]))
                    tempSlice2 := make([]int, len(ruleSlice[secondIDX:firstIDX]))
                    tempSlice3 := make([]int, len(ruleSlice[firstIDX+1:]))
                    copy(tempSlice, ruleSlice[:secondIDX])
                    copy(tempSlice2, ruleSlice[secondIDX:firstIDX])
                    copy(tempSlice3, ruleSlice[firstIDX+1:])
                    //fmt.Println("tempSlice",tempSlice)
                    //fmt.Println("tempSlice2",tempSlice2)
                    //fmt.Println("tempSlice3",tempSlice3)
                    //fmt.Println("ruleSlice BEFORE", ruleSlice)
                    ruleSlice = append(tempSlice, ruleSlice[firstIDX])
                    ruleSlice = append(ruleSlice, tempSlice2[0:]...)
                    ruleSlice = append(ruleSlice, tempSlice3[0:]...)
                    //fmt.Println("ruleSlice AFTER", ruleSlice)
                    break
                }
            }
        }
        fmt.Println(ruleSlice)
        fmt.Println("--------")
    }

    fmt.Println(ruleSlice)

    fileUpd, err := os.Open("./input.updates")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scannerUpd := bufio.NewScanner(fileUpd)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }


    result := 0
    for scannerUpd.Scan() {
        lineUpd := scannerUpd.Text()
        strsUpd := strings.Split(lineUpd, ",")
        currentSlice := make([]int, len(strsUpd))
        for i := range currentSlice {
            currentSlice[i], e = strconv.Atoi(strsUpd[i])
            if e != nil {
                panic(e)
            }
        }

        orderSlice := make([]int, 0)

        //fmt.Println("_____________")
        //fmt.Println(ruleSlice)
        //fmt.Println(currentSlice)

        for i := range ruleSlice {
            for j := range currentSlice {
                if ruleSlice[i] == currentSlice[j] {
                    orderSlice = append(orderSlice, j)
                    //fmt.Println(orderSlice)
                    break
                }
            }
        }

        //fmt.Println("ORDER SLICE LEN",len(orderSlice), orderSlice)
        for i := range len(orderSlice)-1 {
            if orderSlice[i] >= orderSlice[i+1] {
                break
            }
            if (i+1) == len(orderSlice)-1 {
                //fmt.Println(currentSlice[(len(orderSlice)-1)/2])
                result += currentSlice[((len(orderSlice)-1)/2)]
            }
        }
    }

    fmt.Println(result)

}
