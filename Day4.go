package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "regexp"
    "time"
    "sort"
    "strconv"
)

type logEntry struct {
    date    time.Time
    guard   string
    action  string
}

type guardLog struct {
    sleepyTime  int
    popularMins []int
}

var layout = "2006-01-02 15:04"
var linePattern = regexp.MustCompile(`\[(.+)\]\s(.+)`)
var guardIdPattern = regexp.MustCompile(`#(\d+)`)

func getMinArray() []int {
    var log []int
    for i := 0; i <= 59; i++ {
        // Setup a list of minitues to log popularity against
        log = append(log, 0)
    }
    return log
}

func getLines() []string {
    file, err := os.Open("data/day4-input.txt")

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

func part1_and_part2(part string) {
    lines := getLines()

    var logs []logEntry

    // Parse the logs....
    for _, line := range lines {
        matches := linePattern.FindStringSubmatch(line)
        date := matches[1]
        parsedDateTime, _ := time.Parse(layout, date)

        action := matches[2]
        var guard = ""

        if action == "falls asleep" {
            action = "ASLEEP"
        } else if action == "wakes up" {
            action = "WAKE"
        } else {
            // Attempt to get the guard ID
            actionMatches := guardIdPattern.FindStringSubmatch(action)
            action = "BEGIN"
            guard = actionMatches[1]
        }

        var newEntry = logEntry{date: parsedDateTime, action: action, guard: guard}
        logs = append(logs, newEntry)
    }

    sort.Slice(logs, func(i, j int) bool {
        return logs[i].date.Before(logs[j].date)
    })

    var guardLogs = map[string]guardLog{}
    var currentGuard string
    var fellAsleep time.Time
    var sleepMin int

    for _, log := range logs {
        //fmt.Println(log)
        if log.action == "BEGIN" {
            currentGuard = log.guard
        } else if log.action == "ASLEEP" {
            fellAsleep = log.date
            sleepMin = log.date.Minute()
        } else if log.action == "WAKE" {
            wake := log.date
            timeAsleep := wake.Sub(fellAsleep)
            timeAsleepMins := int(timeAsleep.Minutes())

            fmt.Printf("Guard %s been asleept for %d\n", currentGuard, timeAsleepMins)

            if log, ok := guardLogs[currentGuard]; ok {
                log.sleepyTime += timeAsleepMins
                guardLogs[currentGuard] = log
            } else {
                guardLogs[currentGuard] = guardLog{sleepyTime: timeAsleepMins, popularMins: getMinArray()}
            }

            wakeMin := wake.Minute()

            // Mark all the minutes inbetween as asleep
            for sleepMin < wakeMin {
                guardLogs[currentGuard].popularMins[sleepMin] += 1
                sleepMin += 1
            }
        }
    }

    var sleepyGuard string
    var sleepyMins = 0

    for guard, guardLogData := range guardLogs {
        if guardLogData.sleepyTime > sleepyMins {
            sleepyGuard = guard
            sleepyMins = guardLogData.sleepyTime
        }
    }

    fmt.Printf("Sleepy guard is: %s who was sleeping for %d min\n", sleepyGuard, sleepyMins)

    var mostPopularMin = 0
    var mostPopularMinTime = 0

    for min, time := range guardLogs[sleepyGuard].popularMins {
        if time > mostPopularMinTime {
            mostPopularMinTime = time
            mostPopularMin = min
        }
    }

    fmt.Printf("Most popular min: %d with %d asleep mins\n", mostPopularMin, mostPopularMinTime)

    guardId, _ := strconv.Atoi(sleepyGuard)
    checksum :=  guardId * mostPopularMin

    if part == "part1" {
        fmt.Printf("Checksum is: %d\n", checksum)
    } else if part == "part2" {

    }
}

func part2() {

}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        panic("Send part1 or part2")
    }

    part := string(args[0])

    part1_and_part2(part)
}
