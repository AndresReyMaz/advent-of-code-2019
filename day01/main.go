package main

import (
  "fmt"
  "io"
  "math"
  "strconv"
)

func readInput() []int {
  var rawInput string
  var input []int
  for {
    _, err := fmt.Scanln(&rawInput)
    if err == io.EOF {
      break
    }
    i, _ := strconv.Atoi(rawInput)
    input = append(input, i)
  }
  return input
}

func calculateFuel(mass int) int {
  return int(math.Floor(float64(mass) / 3.0)) - 2
}

func part1(input []int) int {
  var ans int
  for _, i := range input {
    ans += calculateFuel(i)
  }
  return ans
}

func part2(input []int) int {
  var ans int
  for _, mass := range input {
    var fuel int
    for {
      currentFuel := calculateFuel(mass)
      if currentFuel <= 0 {
        break
      }
      fuel += currentFuel
      mass = currentFuel
    }
    ans += fuel
  }
  return ans
}

func main() {
  input := readInput()
  fmt.Println(part1(input))
  fmt.Println(part2(input))
}
