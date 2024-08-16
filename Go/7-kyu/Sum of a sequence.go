package kata

func SequenceSum(start, end, step int) int {
  result := 0
  
  if start == end {
    return start
  }
  
  for i := start; i <= end; i += step {
    result += i
  }
  
  return result
}
