package kata

func MoveZeros(arr []int) []int {
  result := make([]int, len(arr))
  
  j := 0
  
  for _, num := range arr {
    if num != 0 {
      result[j] = num
      j++ 
    }
  }
  
  return result
}
