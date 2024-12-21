package main

import (
    "bufio"
    "fmt"
    "log"
    "strconv"
    "os"
    "strings"
    "slices"
)

func main() {

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/2.1/input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    SaveSteps := []int{1,2,3}
    OrderOfReport := "1"
    var SaveReportsCnt int
    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    for scanner.Scan() {
        Line := scanner.Text()
        if err != nil {
           log.Fatalf("unable conver to int", err)
        }

        strings := strings.Split(Line, " ")
        Report := make([]int, len(strings))

        for i, s := range strings {
            Report[i], _ = strconv.Atoi(s)

        }
        ProblemDumpener := 0
        for i := range Report {
            //fmt.Println("Index, len-1, Value", i, len(strings) -2, Report[i])
            if Report[i] > Report[i +1] && i == 0 {
                fmt.Println(i, Report[i], Report[i +1])
                OrderOfReport = "inc"
                fmt.Println(OrderOfReport)
            } else if Report[i] < Report[i +1] && i == 0{
                fmt.Println(i, Report[i], Report[i +1])
                OrderOfReport = "dec"
                fmt.Println(OrderOfReport)
            }


            if OrderOfReport == "inc" {
                if slices.Contains(SaveSteps, (Report[i] - Report[i + 1])) {
                    if i == len(strings) -2 {
                        SaveReportsCnt = SaveReportsCnt + 1
                        fmt.Println("SAVE")
                        break
                    } else {
                    continue
                    }
                } else {
                    if ProblemDumpener != 1 {
                        ProblemDumpener = 1
                        //remove i+1 index from slice
                        copy(Report[i:], Report[i+1:])
                        Report = Report[:len(Report)-1]
                        fmt.Println(Report)
                        //Check new slice
                        if slices.Contains(SaveSteps, (Report[i] - Report[i + 1])) {
                            if i == len(strings) -2 {
                                SaveReportsCnt = SaveReportsCnt + 1
                                fmt.Println("SAVE")
                                break
                            } else {
                            continue
                            }
                            continue
                        }
                    break
                    }
                }
            } else {
                if slices.Contains(SaveSteps, (Report[i + 1] - Report[i])) {
                    if i == len(strings) -2 {
                        SaveReportsCnt = SaveReportsCnt + 1
                        fmt.Println("SAVE")
                        break
                    } else {
                    continue
                    }
                } else {
                    if ProblemDumpener != 1 {
                        ProblemDumpener = 1
                        //remove i+1 index from slice
                        copy(Report[i:], Report[i+1:])
                        Report = Report[:len(Report)-1]
                        fmt.Println(Report)
                        if slices.Contains(SaveSteps, (Report[i] - Report[i + 1])) {
                            if i == len(strings) -2 {
                                SaveReportsCnt = SaveReportsCnt + 1
                                fmt.Println("SAVE")
                                break
                            } else {
                            continue
                            }
                        }
                        continue
                    }
                }
            }
         }
    }
    fmt.Println(SaveReportsCnt)

}
