package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	nMoons = 4
)

type Velocity struct {
	x, y, z int
}

type Moon struct {
	x, y, z int
	vel     Velocity
}

func readInput() []Moon {
	moons := make([]Moon, nMoons)
	for i := 0; i < nMoons; i++ {
		var xString, yString, zString string
		fmt.Scanln(&xString)
		fmt.Scanln(&yString)
		fmt.Scanln(&zString)
		x, _ := strconv.Atoi(xString[3 : len(xString)-1])
		y, _ := strconv.Atoi(yString[1 : len(yString)-1])
		z, _ := strconv.Atoi(zString[1 : len(zString)-1])
		moons[i] = Moon{
			x:   x,
			y:   y,
			z:   z,
			vel: Velocity{x: 0, y: 0, z: 0},
		}
	}
	return moons
}

func applyGravity(moons []Moon) {
	for i := range moons {
		for j := i + 1; j < nMoons; j++ {
			if moons[i].x < moons[j].x {
				moons[i].vel.x++
				moons[j].vel.x--
			} else if moons[i].x > moons[j].x {
				moons[i].vel.x--
				moons[j].vel.x++
			}
			if moons[i].y < moons[j].y {
				moons[i].vel.y++
				moons[j].vel.y--
			} else if moons[i].y > moons[j].y {
				moons[i].vel.y--
				moons[j].vel.y++
			}
			if moons[i].z < moons[j].z {
				moons[i].vel.z++
				moons[j].vel.z--
			} else if moons[i].z > moons[j].z {
				moons[i].vel.z--
				moons[j].vel.z++
			}
		}
	}
}

func applyVelocity(moons []Moon) {
	for i, moon := range moons {
		moons[i].x += moon.vel.x
		moons[i].y += moon.vel.y
		moons[i].z += moon.vel.z
	}
}

func simulateStep(moons []Moon) {
	applyGravity(moons)
	applyVelocity(moons)
}

func simulateSteps(nSteps int, moons []Moon) {
	for i := 1; i <= nSteps; i++ {
		simulateStep(moons)
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func calculateEnergy(moon Moon) int {
	return (abs(moon.x) + abs(moon.y) + abs(moon.z)) * (abs(moon.vel.x) + abs(moon.vel.y) + abs(moon.vel.z))
}

func calculateTotalEnergy(moons []Moon) int {
	energy := 0
	for _, moon := range moons {
		energy += calculateEnergy(moon)
	}
	return energy
}

func part1(moons []Moon) int {
	simulateSteps(1000, moons)
	return calculateTotalEnergy(moons)
}

func part2(moons []Moon) int {
	var xCycle, yCycle, zCycle int
	for step := 1; xCycle == 0 || yCycle == 0 || zCycle == 0; step++ {
		simulateStep(moons)
		var xZero, yZero, zZero int
		for _, moon := range moons {
			if moon.vel.x == 0 {
				xZero++
			}
			if moon.vel.y == 0 {
				yZero++
			}
			if moon.vel.z == 0 {
				zZero++
			}
		}
		if xCycle == 0 && xZero == nMoons {
			xCycle = step
		}
		if yCycle == 0 && yZero == nMoons {
			yCycle = step
		}
		if zCycle == 0 && zZero == nMoons {
			zCycle = step
		}
	}
	cycles := []int{xCycle, yCycle, zCycle}
	var ans int
	sort.Ints(cycles)
	// Iterate using the largest 'cycle'.
	for i := cycles[2]; ; i += cycles[2] {
		// If i is multiple of all three cycles, we found answer.
		if i%cycles[0] == 0 && i%cycles[1] == 0 {
			ans = i
			break
		}
	}
	return ans * 2
}

func main() {
	moons := readInput()
	moonsCpy := make([]Moon, nMoons)
	copy(moonsCpy, moons)
	fmt.Println("Part 1: ", part1(moonsCpy))
	// Warning: takes about ~15sec
	fmt.Println("Part 2: ", part2(moons))
}
