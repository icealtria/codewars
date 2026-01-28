package kata
​
import (
  "strings"
  "unicode"
)
​
func toWeirdCase(str string) string {
  words := []string{}
  for _, w := range strings.Split(str, " ") {
    words = append(words, weirdWord(w))
  }
  return strings.Join(words, " ")
}
​
func weirdWord(w string) string {
  runes := []rune{}
  for i, r := range w {
    if i%2 == 0 {
      runes = append(runes, unicode.ToUpper(r))
    } else {
      runes = append(runes, unicode.ToLower(r))
    }
  }
  return string(runes)
}
​