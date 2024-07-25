package kata

import (
  "strings"
)

func DNAtoRNA(dna string) string {
  result := strings.ReplaceAll(dna, "T", "U")
  
  return result
}
