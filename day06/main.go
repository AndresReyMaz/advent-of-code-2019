package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// Orbit is a space object's orbit.
type Orbit struct {
  name string
  children []*Orbit
}

func readInput() [][]string {
  var input [][]string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    splitInput := strings.Split(scanner.Text(), ")")
    input = append(input, []string{splitInput[0], splitInput[1]})
  }
  return input
}

func buildGraph(input [][]string) *Orbit {
  pointerMap := map[string]*Orbit{}
  for _, rel := range input {
    var from *Orbit
    var to *Orbit
    if orbit, ok := pointerMap[rel[0]]; ok {
      // Orbit has already been created.
      from = orbit
    } else {
      from = &Orbit{rel[0], []*Orbit{}}
      pointerMap[rel[0]] = from
    }
    if orbit, ok := pointerMap[rel[1]]; ok {
      to = orbit
    } else {
      to = &Orbit{rel[1], []*Orbit{}}
      pointerMap[rel[1]] = to
    }
    // Create the edge.
    from.children = append(from.children, to)
  }
  if orbit, ok := pointerMap["COM"]; ok {
    return orbit
  }
  fmt.Println("Center of Mass (COM) not found.")
  return nil
}

func traverseAndSumDistances(orbit *Orbit, dist int) int {
  if orbit == nil {
    return 0
  }
  ans := dist
  for _, child := range orbit.children {
    ans += traverseAndSumDistances(child, dist + 1)
  }
  return ans
}

func hasPathToOrbit(from *Orbit, to string) bool {
  if from.name == to {
    return true
  }
  for _, child := range from.children {
    if hasPathToOrbit(child, to) {
      return true
    }
  }
  return false
}

func findLCA(orbit *Orbit, youOrbit, sanOrbit string) *Orbit {
  hasPathToYouOrbit := hasPathToOrbit(orbit, youOrbit)
  hasPathToSanOrbit := hasPathToOrbit(orbit, sanOrbit)
  if !hasPathToYouOrbit || !hasPathToSanOrbit {
    return nil
  }
  for _, child := range orbit.children {
    if o := findLCA(child, youOrbit, sanOrbit); o != nil {
      return o
    }
  }
  return orbit
}

func getDistanceFrom(orbit *Orbit, name string) int {
  if orbit.name == name {
    return 0
  }
  for _, child := range orbit.children {
    if dist := getDistanceFrom(child, name); dist != -1 {
      return dist + 1
    }
  }
  return -1
}

func part1(input [][]string) int {
  orbitGraph := buildGraph(input)
  return traverseAndSumDistances(orbitGraph, 0)
}

func part2(input [][]string) int {
  // Find the names of the direct orbits for YOU and SAN.
  var youOrbit, sanOrbit string
  for _, rel := range input {
    if rel[1] == "YOU" {
      youOrbit = rel[0]
    } else if rel[1] == "SAN" {
      sanOrbit = rel[0]
    }
  }

  // Build the graph.
  orbitGraph := buildGraph(input)

  // Find the Lowest Common Ancestor.
  lca := findLCA(orbitGraph, youOrbit, sanOrbit)
  youOrbitDist := getDistanceFrom(lca, youOrbit)
  sanOrbitDist := getDistanceFrom(lca, sanOrbit)
  return youOrbitDist + sanOrbitDist
}

func main() {
  input := readInput()
  fmt.Println("Part 1:", part1(input))
  fmt.Println("Part 2:", part2(input))
}