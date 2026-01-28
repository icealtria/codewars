package kata
​
func TwoSum(numbers []int, target int) [2]int {
  memo := make(map[int]int)
  for i, v := range numbers {
    if j, ok := memo[target-v]; ok {
      return [2]int{j, i}
    }
    memo[v] = i
  }
  return [2]int{}
}
​