package tasks

import (
	"fmt"
	"math"
	"strconv"
)

type LuckyTickets struct {
}

func (l LuckyTickets) Run(data []string) string {

	if n, err := strconv.Atoi(data[0]); len(data) > 0 && err == nil {
		return calc(n)
	}

	return ""
}

func calc(n int) string {
	prev := make([]int, 0, 9)

	for i := 0; i < 10; i++ {
		prev = append(prev, 1)
	}

	target := prev

	for i := 2; i <= n; i++ {
		target = make([]int, n*9+1)
		for j := 0; j < len(target); j++ {
			sum := 0

			for k := 9; k >= 0; k-- {
				index := j - k
				if index >= 0 && index < len(prev) {
					sum += prev[index]
				}
			}

			target[j] = sum
		}

		prev = target
	}

	count := float64(0)

	for _, v := range target {
		count += math.Pow(float64(v), float64(2))
	}

	return fmt.Sprintf("%.0f", count)
}
