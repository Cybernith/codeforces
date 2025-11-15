package main

import (
    "bufio"
    "fmt"
    "os"
)

type DisturbedPeopleSolver struct {
    Flats []int
}

func (s DisturbedPeopleSolver) MinimumSwitchOffs() int {
    flats := make([]int, len(s.Flats))
    copy(flats, s.Flats)

    n := len(flats)
    operations := 0

    for i := 1; i+1 < n; i++ {
        if flats[i-1] == 1 && flats[i] == 0 && flats[i+1] == 1 {
            flats[i+1] = 0
            operations++
        }
    }

    return operations
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var n int
    if _, err := fmt.Fscan(in, &n); err != nil {
        return
    }

    flats := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &flats[i])
    }

    solver := DisturbedPeopleSolver{Flats: flats}
    result := solver.MinimumSwitchOffs()
    fmt.Fprintln(out, result)
}
