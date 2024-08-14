package kata

import "strconv"

func SumMix(arr []any) int {
  var sum int
  
  for _, item := range arr {
    switch v := item.(type) {
      case int: 
        sum += v
      case string:
        if num, err := strconv.Atoi(v); err == nil {
          sum += num
        }
    }
  }
  
  return sum
}
