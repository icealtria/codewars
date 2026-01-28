package kata
​
func FindOutlier(integers []int) int {
  lastOdd, lastEven := 0, 0
  oddCount, evenCount := 0, 0
​
  for _, x := range integers {
    if x%2 == 0 {
      lastEven = x
      evenCount++
    } else {
      lastOdd = x
      oddCount++
    }
​
    if oddCount > 0 && evenCount > 0 {
      if oddCount > 1 {
        return lastEven
      }
      if evenCount > 1 {
        return lastOdd
      }
    }
  }
  return 0
}
​