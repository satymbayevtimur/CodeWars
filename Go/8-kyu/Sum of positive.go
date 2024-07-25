package kata

func PositiveSum(numbers []int) int {
  result := 0
  
  for i := 0; i < len(numbers); i++ {
    if numbers[i] < 0 {
      continue
    } else {
      result += numbers[i]
    }
  }
  
  return result
}
