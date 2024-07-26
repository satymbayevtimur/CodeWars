package kata

import (
	"strconv"
	"strings"
)

func countSheep(num int) string {
	if num == 0 {
		return ""
	}

	var builder strings.Builder

	for i := 1; i <= num; i++ {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(" sheep...")
	}

	return builder.String()
}
