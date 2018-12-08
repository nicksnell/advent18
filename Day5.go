package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "io/ioutil"
    "sort"
)

func reducePolymer(polymer string) string {
    // fmt.Println(polymer)

    for index, char := range polymer {
        if index == 0 {
            continue
        }

        lhs := string(char)
        rhs := string(polymer[index - 1])

        if lhs == rhs {
            continue
        } else if strings.ToUpper(lhs) == rhs || lhs == strings.ToUpper(rhs) {
            polymer := polymer[:index - 1] + polymer[index + 1:]
            return reducePolymer(polymer)
        }
    }

    return polymer
}

func getData() string {
    // file, err := ioutil.ReadFile("tests/day5.txt")
    file, err := ioutil.ReadFile("data/day5.txt")

    if err != nil {
        log.Fatal(err)
    }

    polymer := string(file)
    polymer = strings.Replace(polymer, "\n", "", -1)

    return polymer
}

func part1() {
    polymer := getData()
    reducedPolymer := reducePolymer(polymer)
    // fmt.Printf("Reduced Polymer: %s\n", reducedPolymer)
    fmt.Printf("Reduced Polymer Length: %d\n", len(reducedPolymer))
}

func part2() {
    polymer := getData()

    var alphabet = "abcdefghijklmnopqrstuvwxyz"
    var results = map[string]int{}

    for _, char := range alphabet {
        var testPolymer = polymer
        char := string(char)

        testPolymer = strings.Replace(testPolymer, char, "", -1)
        testPolymer = strings.Replace(testPolymer, strings.ToUpper(char), "", -1)
        reducedPolymer := reducePolymer(testPolymer)
        results[char] = len(reducedPolymer)

        fmt.Printf("%s/%s reduces to length %d\n", char, strings.ToUpper(char), len(reducedPolymer))
    }

    var values []int
    for _, v := range results {
        values = append(values, v)
    }

    sort.Ints(values)
    fmt.Printf("Smallest polymer is: %d\n", values[0])
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
