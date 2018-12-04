package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "regexp"
    "strconv"
)

var overlapped = 99999
var areaWidth = 1000 // 10
var areaHeight =  1000 // 8
var pattern = regexp.MustCompile(`#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)

func contains(list []int, search int) bool {
    for _, item := range list {
        if item == search {
            return true
        }
    }
    return false
}

func getLines() []string {
    file, err := os.Open("data/day3-input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string

    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, string(line))
    }

    return lines
}

func getFabricArea(width int, height int) [][]int {
    var area [][]int

    for i := 0; i <= height; i++ {
        var row []int
        for j := 0; j <= width; j++ {
            row = append(row, 0)
        }
        area = append(area, row)
    }

    return area
}

func drawArea(area [][]int) {
    for _, row := range area {
        for _, item := range row {
            if item == overlapped {
                fmt.Print("X")
            } else if item != 0 {
                fmt.Print(item)
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("\n")
    }
}

func parse(claim string) []string {
    // Format is:
    // #1 @ 1,3: 4x4
    // ID, Left Cord/Top Cord, Width/Height
    matches := pattern.FindStringSubmatch(claim)
    return matches
}

func part1_and_part2(part string) {
    lines := getLines()
    area := getFabricArea(areaWidth, areaHeight)

    var ids []int
    var overLappingClaims []int
    var overLappingCount = 0

    for _, line := range lines {
        // Fetch the detail of the claim
        claim := parse(line)

        // Map the claim into the area
        id, _ := strconv.Atoi(claim[1])
        ids = append(ids, id)

        leftCord, _ := strconv.Atoi(claim[2])
        topCord, _ := strconv.Atoi(claim[3])
        width, _ := strconv.Atoi(claim[4])
        height, _ := strconv.Atoi(claim[5])

        for i := 0; i < height; i++ {
            rowIndex := topCord + i
            row := area[rowIndex]

            for j := 0; j < width; j++ {
                colIndex := leftCord + j

                // Check if it's been set before
                if row[colIndex] == 0 {
                    row[colIndex] = id
                } else {
                    // It's not been overlapped before...
                    if row[colIndex] != overlapped {
                        overLappingCount += 1
                    }

                    // Add the overlapping ID
                    if !contains(overLappingClaims, id) {
                        overLappingClaims = append(overLappingClaims, id)
                    }

                    var previousId = row[colIndex]

                    if previousId != overlapped {
                        // If the previous id had not been overlapped before
                        // it wont be in the list, we need to add it also
                        if !contains(overLappingClaims, previousId) {
                            overLappingClaims = append(overLappingClaims, previousId)
                        }
                    }

                    // Mark the area as overlapped...
                    row[colIndex] = overlapped
                }
            }
        }
    }

    if part == "part2" {
        fmt.Printf("Overlapping ID's: %d\n", len(overLappingClaims))

        for _, id := range ids {
            if !contains(overLappingClaims, id) {
                fmt.Printf("Non overlapped claim is: %d\n", id)
            }
        }
    } else {
        // drawArea(area)
        fmt.Printf("Overlapping area: %d\n", overLappingCount)
    }
}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        panic("Send part1 or part2")
    }

    part := string(args[0])

    part1_and_part2(part)
}
