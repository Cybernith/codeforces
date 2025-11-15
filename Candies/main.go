package main

import (
    "bufio"
    "fmt"
    "os"
)

type CandiesDistributor struct {
    TotalCandies int
    FriendsCount int
}

func (c CandiesDistributor) Distribute() []int {
    base := c.TotalCandies / c.FriendsCount
    remainder := c.TotalCandies % c.FriendsCount

    distribution := make([]int, 0, c.FriendsCount)

    for i := 0; i < remainder; i++ {
        distribution = append(distribution, base+1)
    }
    for i := remainder; i < c.FriendsCount; i++ {
        distribution = append(distribution, base)
    }

    return distribution
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var n, m int
    if _, err := fmt.Fscan(in, &n, &m); err != nil {
        return
    }

    distributor := CandiesDistributor{
        TotalCandies: n,
        FriendsCount: m,
    }

    distribution := distributor.Distribute()

    for i, v := range distribution {
        if i > 0 {
            fmt.Fprint(out, " ")
        }
        fmt.Fprint(out, v)
    }
    fmt.Fprintln(out)
}
