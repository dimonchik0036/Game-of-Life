package main

import (
	"os"
	"os/exec"
	"time"
)

const weight = 50
const height = 25

const newLife = 3
const life = 2

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printArray(Array [height][weight]int) {
	for i := -1; i <= height; i++ {
		for j := -1; j <= weight; j++ {
			if i == -1 || i == height || j == -1 || j == weight {
				print("#")
			} else {
				if Array[i][j] == 1 {
					print("О")
				} else {
					print(".")
				}
			}
		}
		println()
	}
}

func nextArray(currentArray [height][weight]int) [height][weight]int {
	var nextArray [height][weight]int
	var count, newu, newl int

	for i := 0; i < height; i++ {
		for j := 0; j < weight; j++ {
			count = 0

			for l := -1; l <= 1; l++ {
				for u := -1; u <= 1; u++ {
					if (l != 0) || (u != 0) {
						newl = 0
						newu = 0

						if i+u < 0 {
							newu += height
						}

						if i+u >= height {
							newu -= height
						}

						if j+l < 0 {
							newl += weight
						}

						if j+l >= weight {
							newl -= weight
						}

						if currentArray[i+u+newu][j+l+newl] == 1 {
							count++
						}
					}
				}
			}

			if (currentArray[i][j] == 0) && (count == newLife) {
				nextArray[i][j] = 1
			} else {
				if (currentArray[i][j] == 1) && (count == newLife || count == life) {
					nextArray[i][j] = 1
				} else {
					nextArray[i][j] = 0
				}
			}
		}
	}

	return nextArray
}

func lifeDrawLine(currentArray [height][weight]int, i int, j int) [height][weight]int {
	currentArray[chekHeight(i)][chekWeight(j)] = 1
	currentArray[chekHeight(i)][chekWeight(j-1)] = 1
	currentArray[chekHeight(i)][chekWeight(j+1)] = 1

	return currentArray
}

func lifeDrawPlaner(currentArray [height][weight]int, i int, j int) [height][weight]int {
	currentArray[chekHeight(i+1)][chekWeight(j-1)] = 1
	currentArray[chekHeight(i+1)][chekWeight(j)] = 1
	currentArray[chekHeight(i+1)][chekWeight(j+1)] = 1
	currentArray[chekHeight(i)][chekWeight(j+1)] = 1
	currentArray[chekHeight(i-1)][chekWeight(j)] = 1

	return currentArray
}

func chekHeight(i int) int {
	var newI int

	for i+newI < 0 {
		newI += height
	}

	for i+newI >= height {
		newI -= height
	}

	return i + newI
}

func chekWeight(j int) int {
	var newJ int

	for j+newJ < 0 {
		newJ += weight
	}

	for j+newJ >= weight {
		newJ -= weight
	}

	return j + newJ
}

func main() {
	clear()

	currentArray := [height][weight]int{}

	currentArray = lifeDrawLine(currentArray, 10, 17)
	currentArray = lifeDrawPlaner(currentArray, 10, 21)

	printArray(currentArray)

	for {
		time.Sleep(40 * time.Millisecond)
		clear()
		currentArray = nextArray(currentArray)
		printArray(currentArray)
	}
}
