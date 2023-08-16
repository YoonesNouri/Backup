package main

import (
    "fmt"
    "strconv"
)

func main() {
    var numStr string
    fmt.Print("Enter a number: ")
    fmt.Scanln(&numStr)

    num, err := strconv.Atoi(numStr)
    if err != nil {
        fmt.Println("Invalid number entered.")
        return
    }

    digits := countDigits(num)
    fmt.Printf("Number of digits: %d\n", digits)
}

func countDigits(num int) int {
    if num == 0 {
        return 1
    }

    count := 0
    for num != 0 {
        num /= 10
        count++
    }

    return count
}

