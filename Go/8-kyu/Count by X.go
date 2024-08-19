package kata

func CountBy(x, n int) []int {
    result := make([]int, n)
  
    for i := 0; i < n; i++ {
        result[i] = x * (i + 1)
    }
  
    return result
}
