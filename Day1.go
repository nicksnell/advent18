package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func findFrequency(frequency int) int {
    file, err := os.Open("data/day1-input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        sign := line[0]
        adjust, _ := strconv.Atoi(string(line[1:]))

        if string(sign) == "-" {
            adjust = adjust * -1
        }

        frequency += adjust
    }

    return frequency
}

func contains(list []int, search int) bool {
    for _, item := range list {
        if item == search {
            return true
        }
    }
    return false
}

func part1() {
    result := findFrequency(0)
    fmt.Println(result)
}

func part2() {
    // TODO: Refactor
    var knownFrequencies []int
    var repeatedFrequency int
    var searchForFrequency = true
    var frequency = 0

    for searchForFrequency {
        file, err := os.Open("data/day1-input.txt")

        if err != nil {
            log.Fatal(err)
        }

        defer file.Close()

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            line := scanner.Text()
            sign := line[0]
            adjust, _ := strconv.Atoi(string(line[1:]))

            if string(sign) == "-" {
                adjust = adjust * -1
            }

            frequency += adjust

            if contains(knownFrequencies, frequency) {
                repeatedFrequency = frequency
                searchForFrequency = false
                break
            }

            knownFrequencies = append(knownFrequencies, frequency)
        }
    }

    fmt.Println(repeatedFrequency)
}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        panic("Send part1 or part2")
    }

    part := string(args[0])

    if part == "part1" {
        part1()
    } else if part == "part2" {
        part2()
    }
}
