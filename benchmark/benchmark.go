package benchmark

import (
	"fmt"
	"time"
)

func BenchmarkFunction(fn func(), count int, rounds int) {
	var totalDuration time.Duration

	for r := 0; r < rounds; r++ {
		start := time.Now()

		for i := 0; i < count; i++ {
			fn()
		}

		duration := time.Since(start)
		totalDuration += duration
	}

	averageDuration := totalDuration / time.Duration(rounds)
	fmt.Printf("\nAverage time over %d rounds: %v\n", rounds, averageDuration)
	fmt.Printf("Average time per operation: %v\n", averageDuration/time.Duration(count))
}
