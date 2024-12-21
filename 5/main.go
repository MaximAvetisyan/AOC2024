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

func orderUpdates(updates[]int, ruleMap map[int][]int) []int {
    result := make([]int, 0)
    indexMap := make(map[int]int)
    for i , data := range updates {
        newIndex := len(result)
        for d := range result {
            indexMap[result[d]] = d
        }
        v, ok := ruleMap[updates[i]]
        if ok {
            for k := range v {
                if idx, ok := indexMap[v[k]]; ok {
                    if newIndex > idx {
                        newIndex = idx
                    }
                }
            }
        }
        endSlice := slices.Clone(result[:newIndex])
        result = append(result, data)
        result = append(result, endSlice...)
        fmt.Println(endSlice)
        fmt.Println(result)
        fmt.Println("============")
    }

    return result
}

func main () {

    file, err := os.Open("./input.order")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    nums := make([]int,2)
    var e error
    ruleMap := make(map[int][]int,0)

    for scanner.Scan() {
        line := scanner.Text()
        strs := strings.Split(line, "|")
        for i := range nums {
            nums[i], e = strconv.Atoi(strs[i])
            if e != nil {
                panic(e)
            }
        }

        v, ok := ruleMap[nums[0]]
        if !ok {
            ruleMap[nums[0]] = []int{nums[1]}
        } else {
            v = append(v, nums[1])
            ruleMap[nums[0]] = v
        }
    }
    //fmt.Println(ruleMap)

    fileUpd, err := os.Open("./input.updates")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }
    scannerUpd := bufio.NewScanner(fileUpd)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    res := 0
    res2 := 0
    fmt.Println(ruleMap)

    for scannerUpd.Scan() {
        line := scannerUpd.Text()
        strs := strings.Split(line, ",")
        updates := make([]int, len(strs))
        for i := range updates {
            updates[i], e = strconv.Atoi(strs[i])
            if e != nil {
                panic(e)
            }
        }

        //fmt.Println(updates)
        orderIsCorrect := true
        for i := range updates {
            if i == 0 {
                continue
            }
            v, ok := ruleMap[updates[i]]
            if !ok {
                continue
            }
            for k := range v {
                for j := 0; j < i; j++ {
                    if v[k] == updates[j] {
                        orderIsCorrect = false
                    }
                }
            }
        }
        if orderIsCorrect {
            res = res + updates[len(updates)/2]
        } else {
            orderedUpdate := orderUpdates(updates, ruleMap)
            fmt.Println(orderedUpdate)
            fmt.Println("================================")
            res2 = res2 + orderedUpdate[len(orderedUpdate)/2]
        }
    }
    fmt.Println(res)
    fmt.Println(res2)
}
