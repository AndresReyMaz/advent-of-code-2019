package main

import (
  "fmt"
  "strconv"
  "strings"
)

func readInput() (int, int) {
	var input string
	fmt.Scanln(&input)
	splitInput := strings.Split(input, "-")
  begin, _ := strconv.Atoi(splitInput[0])
	end, _ := strconv.Atoi(splitInput[1])
	return begin, end
}

func hasDouble(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
    if s[i] == s[i + 1] {
      return true
    }
  }
  return false
}

func isNondecreasing(s string) bool {
  prev := 0
  for i := 0; i < len(s); i++ {
    intVal, _ := strconv.Atoi(string(s[i]))
    if intVal < prev {
      return false
    }
    prev = intVal
  }
  return true
}

func hasExactDouble(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
    if s[i] == s[i + 1] {
      if i == (len(s) - 2) || s[i] != s[i + 2] {
        return true
      }
      x := i
      for ; x < len(s) && s[i] == s[x]; x++ {
      }
      i = x - 1
    }
  }
  return false
}

func part1(begin, end int) int {
	ans := 0
	for i := begin; i <= end; i++ {
    s := strconv.Itoa(i)
    if hasDouble(s) && isNondecreasing(s) {
      ans++
    }
	}
	return ans
}

func part2(begin, end int) int {
  ans := 0
	for i := begin; i <= end; i++ {
    s := strconv.Itoa(i)
    if hasExactDouble(s) && isNondecreasing(s) {
      ans++
    }
	}
	return ans
}

func main() {
	begin, end := readInput()
	fmt.Println(part1(begin, end))
	fmt.Println(part2(begin, end))
}