package kata

func Summation(n int) int {
  result := 0
  
  for i := n; i >= 1; i-- {
    result += i;
  }
  
  return result
}
