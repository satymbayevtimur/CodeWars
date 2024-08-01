package kata

import (
	"fmt"
)

func SeriesSum(n int) string {
	if n == 0 {
		return "0.00"
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		term := 1.0 / float64(1 + 3 * i)
		sum += term
	}

	roundedSum := fmt.Sprintf("%.2f", sum)
	return roundedSum
}
