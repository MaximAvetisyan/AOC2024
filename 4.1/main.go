package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
)

//func traverseMatrix(textMatrix [][]string, row int, col int) (Found int) {
//
//    var l,d int
//    l = len(textMatrix)
//    if l > 0 {
//        d = len(textMatrix[0])
//    }
//	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
//	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}
//
//    lookupMatrix := []string{"M","A","S"}
//
//    for i := range x {
//        curRow := row + x[i]
//        curCol := col + y[i]
//
//        for k := range lookupMatrix {
//            //Out of bount check
//            if curRow < 0 || curCol < 0 || curRow >= l || curCol >= d {
//                break
//            }
//
//            if textMatrix[curRow][curCol] != lookupMatrix[k] {
//                break
//            }
//
//            //fmt.Println(textMatrix[curRow][curCol])
//
//            curRow += x[i]
//            curCol += y[i]
//
//            if k == len(lookupMatrix) -1 {
//                Found++
//                //fmt.Println(Found)
//            }
//        }
//    }
//
//    return Found
//}

func main () {

    file, err := os.Open("./input.txt")
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    scanner := bufio.NewScanner(file)
    if err != nil {
       log.Fatalf("unable to read file: %v", err)
    }

    textMatrix := make([][]string, 0)

    for scanner.Scan() {
        line := scanner.Text()
        rowSlice := strings.Split(line, "")
        textMatrix = append(textMatrix, rowSlice)
    }

    fmt.Println(textMatrix)
    var TotalRes int

    for row := range len(textMatrix) - 2 {
        for col := range len(textMatrix[row]) - 2 {
            if (textMatrix[row][col] == "M" || textMatrix[row][col] == "S") &&
               (textMatrix[row][col + 2] == "M" || textMatrix[row][col + 2] == "S") &&
               (textMatrix[row+1][col+1] == "A") &&
               ((textMatrix[row+2][col] == "M" || textMatrix[row+2][col] == "S") && textMatrix[row+2][col] != textMatrix[row][col+2]) &&
               ((textMatrix[row+2][col+2] == "M" || textMatrix[row+2][col+2] == "S") && textMatrix[row+2][col+2] != textMatrix[row][col]){
                TotalRes++
                fmt.Println(textMatrix[row][col])
                fmt.Println(textMatrix[row][col+2])
                fmt.Println(textMatrix[row+1][col+1])
                fmt.Println(textMatrix[row+2][col])
                fmt.Println(textMatrix[row+2][col+2])
                fmt.Println(row,col)
            } else {
                continue
            }
        }
    }
    fmt.Println(TotalRes)
}
