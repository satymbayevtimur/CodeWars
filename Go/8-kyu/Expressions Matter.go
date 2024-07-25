package kata

func max(values ...int) int {
  maxValue := values[0]
  
  for _, value := range values {
    if value > maxValue {
      maxValue = value
    }
  }
  
  return maxValue
}

func ExpressionMatter(a int, b int, c int) int {
  firstExp := a * (b + c)
  secondExp := a * b * c
  thirdExp := a + b * c
  fourthExp := (a + b) * c  
  fifthExp := a + b + c
  
  return max(firstExp, secondExp, thirdExp, fourthExp, fifthExp)
}
