package main

import (
	"bytes"
	"time"
)

const weight = 50
const height = 25

const newLife = 3
const life = 2

func nextArray(currentArray [height][weight]bool) [height][weight]bool {
	var nextArray [height][weight]bool
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

						if currentArray[i+u+newu][j+l+newl] == true {
							count++
						}
					}
				}
			}

			if (currentArray[i][j] == false) && (count == newLife) {
				nextArray[i][j] = true
			} else {
				if (currentArray[i][j] == true) && (count == newLife || count == life) {
					nextArray[i][j] = true
				} else {
					nextArray[i][j] = false
				}
			}
		}
	}

	return nextArray
}

func lifeDrawLine(currentArray [height][weight]bool, i int, j int) [height][weight]bool {
	currentArray[chekHeight(i)][chekWeight(j)] = true
	currentArray[chekHeight(i)][chekWeight(j-1)] = true
	currentArray[chekHeight(i)][chekWeight(j+1)] = true

	return currentArray
}

func lifeDrawPlaner(currentArray [height][weight]bool, i int, j int, pos int) [height][weight]bool {
	currentArray[chekHeight(i+1)][chekWeight(j-1)] = true
	currentArray[chekHeight(i+1)][chekWeight(j)] = true
	currentArray[chekHeight(i+1)][chekWeight(j+1)] = true
	currentArray[chekHeight(i-1)][chekWeight(j)] = true

	switch pos {
	case 1:
		currentArray[chekHeight(i)][chekWeight(j-1)] = true
	default:
		currentArray[chekHeight(i)][chekWeight(j+1)] = true
	}

	return currentArray
}

func lifeDrawShip(currentArray [height][weight]bool, i int, j int, pos int) [height][weight]bool {
	currentArray[chekHeight(i+3)][chekWeight(j+1)] = true
    currentArray[chekHeight(i+3)][chekWeight(j+2)] = true
    currentArray[chekHeight(i+3)][chekWeight(j+3)] = true

    var ru, r int
	switch pos {
	case 1:
		ru = -1
        r = -2
	default:
		ru = 1
        r = 2
	}

    currentArray[chekHeight(i)][chekWeight(j+2+ru)] = true
    currentArray[chekHeight(i+1)][chekWeight(j+2+r)] = true
    currentArray[chekHeight(i+2)][chekWeight(j+4)] = true
    currentArray[chekHeight(i+2)][chekWeight(j)] = true
    currentArray[chekHeight(i+3)][chekWeight(j+2+r)] = true

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

func arrayToString(array [height][weight]bool) string {
	var buffer bytes.Buffer

	for i := 0; i < 100; i++ {
		buffer.WriteByte('\n')
	}

	for i := -1; i <= height; i++ {
		for j := -1; j <= weight; j++ {
			b := byte(' ')

			if i == -1 || i == height || j == -1 || j == weight {
				b = '@'
			} else {
				if array[i][j] == true {
					b = '*'
				}
			}

			buffer.WriteByte(b)
		}
		buffer.WriteByte('\n')
	}

	return buffer.String()
}

func main() {

	currentArray := [height][weight]bool{}

	currentArray = lifeDrawLine(currentArray, height/2 - 1, weight/2 + 10)
	currentArray = lifeDrawPlaner(currentArray, 20, 15, 0)
	//currentArray = lifeDrawPlaner(currentArray, 10, 50, 1)
    	currentArray = lifeDrawShip(currentArray, 0, 0, 0)
   	 currentArray = lifeDrawShip(currentArray, 15, 0, 1)

	for {
		print(arrayToString(currentArray))

		currentArray = nextArray(currentArray)
		time.Sleep(70 * time.Millisecond)
	}
}
