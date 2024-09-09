package kata

import "fmt"

func Digitize(n int) []int {
    numStr := fmt.Sprint(n)
    
    var result []int
    
    for i := len(numStr) - 1; i >= 0; i-- {
        digit := numStr[i] - '0'
        result = append(result, int(digit))
    }
    
    return result
}
