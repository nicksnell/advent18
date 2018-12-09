package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
    "sort"
)

type location struct {
    id          int
    x           int
    y           int
    count       int
    infinite    bool
}

type manhatten struct {
    id          int
    distance    int
}

func abs(value int) int {
    if value < 0 {
        value = value * -1
    }
    return value
}

func getLines() []string {
    file, err := os.Open("data/day6.txt")
    // file, err := os.Open("tests/day6.txt")

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

func getGrid(maxX int, maxY int) [][]int {
    var grid [][]int

    for i := 0; i <= maxY; i++ {
        var row []int
        for j := 0; j <= maxX; j++ {
            row = append(row, 0)
        }
        grid = append(grid, row)
    }

    return grid
}

func getLocationsAndGrid() (map[string]location, [][]int, []int) {
    points := getLines()

    // Maximum bounds
    var maxX = 0
    var maxY = 0
    var locations = make(map[string]location)

    // Parse and define the extent of the grid
    for index, point := range points {
        id := index + 1
        point := string(point)
        coords := strings.Split(point, ", ")
        x, _ := strconv.Atoi(coords[0])
        y, _ := strconv.Atoi(coords[1])

        if x > maxX {
            maxX = x
        }

        if y > maxY {
            maxY = y
        }

        // Note: Don't use 0 as an ID as we use this later to
        // check if its been processed
        var location = location{id: id, x: x, y: y, count: 0, infinite: false}
        locations[strconv.Itoa(id)] = location
    }

    // Build a grid to store the results
    grid := getGrid(maxX, maxY)

    // Mark locations into the grid
    for _, location := range locations {
        x, y := location.x, location.y
        grid[y][x] = 100 + location.id
    }

    var extent = []int{maxX, maxY}

    return locations, grid, extent
}

func drawGrid(grid [][]int) {
    // Visual grid for debugging
    for _, row := range grid {
        for _, col := range row {
            fmt.Printf("%d, ", col)
        }
        fmt.Printf("\n")
    }
}

func part1() {
    locations, grid, extent := getLocationsAndGrid()

    var equalDistance = 99999

    maxX := extent[0]
    maxY := extent[1]

    // For each space on the grid, compute the Manhattan distance
    // to each location. The nearest location should have it's ID
    // marked in the grid, and recored. Equidistant spaces are note
    // counted. If a space touches the 'edge' of the grid, it's considered
    // to be an infinite range
    for y, row := range grid {
        for x, col := range row {
            // Check space is not a known location
            if col != 0 {
                continue
            }

            var distances []manhatten

            // Compute each location
            for _, location := range locations {
                xOffset := x - location.x
                yOffset := y - location.y

                manhattenDistance := abs(xOffset) + abs(yOffset)

                var result = manhatten{id: location.id, distance: abs(manhattenDistance)}
                distances = append(distances, result)
            }

            // Compute the closest location & save
            sort.Slice(distances, func(i, j int) bool {
                return distances[i].distance < distances[j].distance
            })

            closest := distances[0]
            secondClosest := distances[1]

            if closest.distance == secondClosest.distance {
                // Space is equidistant
                grid[y][x] = equalDistance
            } else {
                grid[y][x] = closest.id

                var location = locations[strconv.Itoa(closest.id)]

                // Update the location closest count
                location.count += 1

                // If the x or y represent an 'edge'
                // mark the location as having infinate bounds
                if x == 0 || x == maxX {
                    location.infinite = true
                } else if y == 0 || y == maxY {
                    location.infinite = true
                }

                locations[strconv.Itoa(closest.id)] = location
            }
        }
    }

    var largestAreaId int
    var largestArea = 0
    for _, location := range locations {
        if location.infinite == false {
            if location.count > largestArea {
                largestAreaId = location.id
                largestArea = location.count
            }
        }
    }

    fmt.Printf("Largest area is #%d with %d close spaces\n", largestAreaId, (largestArea + 1))

    // drawGrid(grid)
}

func part2() {
    locations, grid, _ := getLocationsAndGrid()

    //var maximumManhattenDistance = 32
    var maximumManhattenDistance = 10000

    var safeRegion = 88888
    var safeRegionSize = 0

    for y, row := range grid {
        for x, _ := range row {
            var totalManhattenDistance = 0

            for _, location := range locations {
                xOffset := x - location.x
                yOffset := y - location.y
                manhattenDistance := abs(xOffset) + abs(yOffset)

                totalManhattenDistance += manhattenDistance
            }

            if totalManhattenDistance < maximumManhattenDistance {
                safeRegionSize += 1

                // Update the grid for visual representation
                grid[y][x] = safeRegion
            }
        }
    }

    fmt.Printf("Safe region size is: %d\n", safeRegionSize)
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
