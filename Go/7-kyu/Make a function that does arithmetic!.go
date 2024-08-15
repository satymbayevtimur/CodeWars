package kata

func Arithmetic(a int, b int, operator string) int {
  result := 0
  
  switch {
    case operator == "add":
      result += a + b
      break
    case operator == "subtract":
      result += a - b
      break
    case operator == "multiply":
      result += a * b
      break
    case operator == "divide":
      result += a / b
      break
  }
  
  return result
}
