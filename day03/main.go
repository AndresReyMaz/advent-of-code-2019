package main

import (
  "fmt"
  "math"
  "strconv"
  "strings"
)

// PathItem contains the magnitude of a wire along a single direction.
type PathItem struct {
  direction string
  magnitude int
}

// Coord is a simple x-y plane coordinate.
type Coord struct {
  x int
  y int
}

// Wire is the collection of PathItems that make a wire.
type Wire = []PathItem

func readWire() Wire {
  var input string
  fmt.Scanln(&input)
  splitInput := strings.Split(input, ",")
  wire := make([]PathItem, len(splitInput))
  for i, location := range splitInput {
    magnitude, _ := strconv.Atoi(location[1:])
    wire[i] = PathItem{string(location[0]), magnitude}
  }
  return wire
}

func readInput() []Wire {
  input := make([]Wire, 2)
  input[0], input[1] = readWire(), readWire()
  return input
}

func moveDirection(dir string, x, y int) (int, int) {
  switch dir {
  case "U":
    y++
  case "D":
    y--
  case "R":
    x++
  case "L":
    x--
  }
  return x, y
}

func generateWireMap(wire Wire) map[Coord]int {
  x := 0
  y := 0
  wireMap := map[Coord]int{}
  step := 0
  for _, item := range wire {
    for i := 0; i < item.magnitude; i++ {
      step++
      x, y = moveDirection(item.direction, x, y)
      if _, ok := wireMap[Coord{x, y}]; !ok {
        // This coordinate has not been visited.
        wireMap[Coord{x, y}] = step
      }
    }
  }
  return wireMap
}

func manhattan(x, y int) int {
  // The central port is at coordinate (0, 0).
  return abs(x) + abs(y)
}

func abs(x int) int {
  if x < 0 {
    return x * -1
  }
  return x
}

func min(x, y int) int {
  if x > y {
    return y
  }
  return x
}

func solve(wires []Wire) (int, int) {
  // Generate a map of coordinates for the first wire.
  wireMap := generateWireMap(wires[0])

  // Traverse the second wire and find the correct intersections.
  x := 0
  y := 0
  closest := math.MaxInt32
  fewest := math.MaxInt32
  step := 0
  
  for _, item := range wires[1] {
    for i := 0; i < item.magnitude; i++ {
      step++
      x, y = moveDirection(item.direction, x, y)
      if _, ok := wireMap[Coord{x, y}]; ok {
        closest = min(closest, manhattan(x, y))
      }

      if firstWireSteps, ok := wireMap[Coord{x, y}]; ok {
        fewest = min(fewest, firstWireSteps + step)
      }
    }
  }
  return closest, fewest
}

func main() {
  input := readInput()
  closest, fewest := solve(input)
  fmt.Println("Part 1: ", closest)
  fmt.Println("Part 2: ", fewest)
}