package kata


func FakeBin(x string) string {
  result := []byte(x)
  
  for i := 0; i < len(result); i++ {
    if result[i] < '5' {
      result[i] = '0'
    } else {
      result[i] = '1'
    }
  }
  
  return string(result)
}
