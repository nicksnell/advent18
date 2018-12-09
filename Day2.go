package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "reflect"
    "strings"
)

func contains(list []reflect.Value, search string) bool {
    for _, item := range list {
        if item.Interface() == search {
            return true
        }
    }
    return false
}

func getLines() []string {
    file, err := os.Open("data/day2.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string

    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    return lines
}

func getLetterMatrix() map[string]map[string]int {
    lines := getLines()

    var matrix = map[string]map[string]int{}

    for _, line := range lines {
        var letters = map[string]int{}

        // Go through each letter and count...
        for _, letter := range line {
            keys := reflect.ValueOf(letters).MapKeys()
            stringLetter := string(letter)

            if contains(keys, stringLetter) {
                letters[stringLetter] += 1
            } else {
                letters[stringLetter] = 1
            }
        }

        matrix[string(line)] = letters
    }

    return matrix
}

func part1() {
    matrix := getLetterMatrix()

    var twoLetterCount = 0
    var threeLetterCount = 0

    for id, letters := range matrix {
        // Calculate any counts
        var foundTwo = false
        var foundThree = false

        for _, count := range letters {
            if count == 2 && !foundTwo {
                foundTwo = true
                twoLetterCount += 1
            }

            if count == 3 && !foundThree {
                foundThree = true
                threeLetterCount += 1
            }
        }

        fmt.Printf("%s : 2 == %t : 3 == %t\n", id, foundTwo, foundThree)
    }

    fmt.Printf("Box IDs with two: %d\n", twoLetterCount)
    fmt.Printf("Box IDs with three: %d\n", threeLetterCount)

    checksum := twoLetterCount * threeLetterCount

    fmt.Printf("Checksum is: %d\n", checksum)
}

func part2() {
    lines := getLines()

    var smallestDiff = 27
    var overallDiffs = map[string]map[string]int{}

    for _, line := range lines {
        var differences = map[string]int{}

        for _, otherLine := range lines {
            if line == otherLine {
                continue
            }

            var difference = 0

            for position, letterRune := range line {
                letter := string(letterRune)
                otherLetter := string(otherLine[position])

                if letter != otherLetter {
                    difference++
                }
            }

            differences[otherLine] = difference

            if difference < smallestDiff {
                smallestDiff = difference
            }
        }

        overallDiffs[line] = differences
    }

    var smallestDiffIds []string

    for id, differences := range overallDiffs {
        for _, diff := range differences {
            if diff == smallestDiff {
                smallestDiffIds = append(smallestDiffIds, id)
            }
        }
    }

    if len(smallestDiffIds) > 2 {
        log.Fatal("Found more than 2 matches!",)
    }

    var key []string
    lhs := smallestDiffIds[0]
    rhs := smallestDiffIds[1]

    for _, letterRune := range lhs {
        letter := string(letterRune)
        if strings.ContainsAny(rhs, letter) {
            key = append(key, letter)
        }
    }

    fmt.Printf("Key: %s\n", strings.Join(key, ""))
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
