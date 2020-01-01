package main

import (
  "fmt"
  "strconv"
  "strings"
)

const (
  _Add = 1
  _Mult = 2
  _Halt = 99
  _DesiredOutput = 19690720
)

func readInput() []int {
  var input string
  fmt.Scanln(&input)
  splitInput := strings.Split(input, ",")
  program := make([]int, len(splitInput))
  for i, x := range splitInput {
    intVal, _ := strconv.Atoi(x)
    program[i] = intVal
  }
  return program
}

func run(program []int, noun int, verb int) int {
  pos := 0
  program[1] = noun
  program[2] = verb
OuterLoop:
  for {
    if pos > len(program) {
      return -1
    }
    switch program[pos] {
    case _Add:
      program[program[pos + 3]] = program[program[pos + 1]] + program[program[pos + 2]]
    case _Mult:
      program[program[pos + 3]] = program[program[pos + 1]] * program[program[pos + 2]]
    case _Halt:
      break OuterLoop
    }
    pos += 4
  }
  return program[0]
}

func part1(program []int) int {
  programCopy := make([]int, len(program))
  copy(programCopy, program)
  return run(programCopy, 12, 2)
}

func part2(program []int) int {
  for noun := 0; noun <= 99; noun++ {
    for verb := 0; verb <= 99; verb++ {
      programCopy := make([]int, len(program))
      copy(programCopy, program)
      output := run(programCopy, noun, verb)
      if output == _DesiredOutput {
        return 100 * noun + verb
      }
    }
  }
  return -1
}

func main() {
  input := readInput()
  fmt.Println(part1(input))
  fmt.Println(part2(input))
}