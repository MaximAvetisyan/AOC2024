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

func RemoveIndex (slice []int, i int) ([]int) {
    ret := make([]int,0)
    ret = append(ret, slice[:i]...)
    ret = append(ret, slice[i+1:]...)
    return ret
}

func CheckReport(report []int, DumpnerCC int) (IsSave int) {

    OrderOfReport := "1"
    SaveSteps := []int{1,2,3}

    for i := range report {
        if report[i] > report[i +1] && i == 0 {
            fmt.Println(report)
            OrderOfReport = "dec"
            fmt.Println(OrderOfReport)
        } else if report[i] < report[i +1] && i == 0{
            fmt.Println(report)
            OrderOfReport = "inc"
            fmt.Println(OrderOfReport)
        }

        if OrderOfReport == "dec" {
            if slices.Contains(SaveSteps, (report[i] - report[i + 1])) {
                if i == (len(report) - 2) {
                    IsSave = 1
                    break
                } else {
                    continue
                }
            } else {
                if DumpnerCC == 1 {
                break

                } else {
                    fmt.Println("FIX ME PLS")
                    for j := range report {
                        reportTry := RemoveIndex(report, j)
                        IsSave = CheckReport(reportTry, 1)
                        if IsSave == 1 {
                            break
                        }
                    }
                    break
                }
            }
        } else {
            if slices.Contains(SaveSteps, (report[i + 1] - report[i])) {
                if i == (len(report) - 2) {
                    IsSave = 1
                    break
                } else {
                    continue
                }
            } else {
                if DumpnerCC == 1 {
                break
                } else {
                    fmt.Println("FIX ME PLS")
                    for j := range report {
                        reportTry := RemoveIndex(report, j)
                        IsSave = CheckReport(reportTry, 1)
                        if IsSave == 1 {
                            break
                        }
                    }
                    break
                }
            }
        }
    }

    if IsSave == 1 {
        fmt.Println("SAVE")
    } else {
        fmt.Println("BROKEN")
    }

    return IsSave
}

func main() {

    file, err := os.Open("/Users/maksimavetisan/Documents/PERSONAL_PROJECTS/AOC2024/2.2/input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

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
        report := make([]int, len(strings))

        for i, s := range strings {
            report[i], _ = strconv.Atoi(s)

        }

        ProblemDumpener := 0
        IsSave := CheckReport(report, ProblemDumpener)
        SaveReportsCnt = SaveReportsCnt + IsSave
   }

    fmt.Println(SaveReportsCnt)

}


