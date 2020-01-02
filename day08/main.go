package main

import (
  "fmt"
  "math"
)

const (
  // Width is the width of a layer.
  Width = 25
  // Height is the height of a layer.
  Height = 6
  // LayerSize is the size of a layer in pixels (25 wide, 6 tall).
  LayerSize = Width * Height
)

func getLayerWithFewest0Digits(data string) int {
  ans := 0
  fewest := math.MaxInt32
  for layer := 0; layer * LayerSize < len(data); layer++ {
    numZeros := 0
    for i := layer * LayerSize; i < (layer + 1) * LayerSize; i++ {
      if string(data[i]) == "0" {
        numZeros++
      }
    }
    if numZeros < fewest {
      fewest = numZeros
      ans = layer
    }
  }
  return ans
}

func countLayerDigits(layer int, data, digit string) int {
  ans := 0
  for i := layer * LayerSize; i < (layer + 1) * LayerSize; i++ {
    if string(data[i]) == digit {
      ans++
    }
  }
  return ans
}

func part1(data string) int {
  layer := getLayerWithFewest0Digits(data)
  return countLayerDigits(layer, data, "1") * countLayerDigits(layer, data, "2")
}

func part2(data string) {
  for i := 0; i < Height; i++ {
    for j := 0; j < Width; j++ {
      for layer := 0; layer * LayerSize < len(data); layer++ {
        pixel := string(data[layer * LayerSize + (i * Width) + j])
        if pixel != "2" {
          fmt.Printf(pixel)
          break
        }
      }
    }
    fmt.Println()
  }
}

func main() {
  var input string
  fmt.Scanln(&input)
  fmt.Println("Part 1:", part1(input))
  fmt.Println("Part 2:")
  part2(input)
}