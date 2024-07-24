package kata


func Invert(arr []int) []int {
  result := make([]int, len(arr))
  
  for i := 0; i < len(arr); i++ {
    result[i] = arr[i] * -1
  }
  
  return result
}
