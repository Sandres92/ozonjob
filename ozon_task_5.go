package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countMatrices int
	fmt.Fscanf(in, "%d\n", &countMatrices)

	inputMatices(countMatrices)
}

func inputMatices(countMatrices int) [][][]rune {
	resultMatrces := make([][][]rune, countMatrices)

	for k := 0; k < countMatrices; k++ {
		var robotA, robotB Point
		resultMatrces[k], robotA, robotB = inputOneMatrix()

		if robotA.x < robotB.x {
			//ведём робота а наверх
			resultMatrces[k] = moveRobots(resultMatrces[k], robotA, 'a', robotB, 'b')
		} else if robotB.x < robotA.x {
			resultMatrces[k] = moveRobots(resultMatrces[k], robotB, 'b', robotA, 'a')
		} else {
			if robotA.y < robotB.y {
				resultMatrces[k] = moveRobots(resultMatrces[k], robotA, 'a', robotB, 'b')
			} else {
				resultMatrces[k] = moveRobots(resultMatrces[k], robotB, 'b', robotA, 'a')
			}
		}
	}

	for k := 0; k < countMatrices; k++ {
		for i := 0; i < len(resultMatrces[k]); i++ {
			for j := 0; j < len(resultMatrces[k][i]); j++ {
				fmt.Fprintf(out, "%c", resultMatrces[k][i][j])
			}
			fmt.Fprintln(out)
		}
	}
	return resultMatrces
}

func inputOneMatrix() ([][]rune, Point, Point) {
	var sizeN, sizeM int
	_, err := fmt.Fscanf(in, "%d %d\n", &sizeN, &sizeM)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	originalMatrix := make([][]rune, sizeN)
	resultMatrix := make([][]rune, sizeN)
	var robotA, robotB Point

	for i := 0; i < sizeN; i++ {

		originalMatrix[i] = make([]rune, sizeM)
		resultMatrix[i] = make([]rune, sizeM)

		for j := 0; j < sizeM; j++ {
			fmt.Fscanf(in, "%c", &originalMatrix[i][j])

			if originalMatrix[i][j] == 'A' {
				robotA.x = i
				robotA.y = j
			} else if originalMatrix[i][j] == 'B' {
				robotB.x = i
				robotB.y = j
			}
			resultMatrix[i][j] = originalMatrix[i][j]
		}
		in.ReadString('\n')
	}

	return originalMatrix, robotA, robotB
}

func isCanMove(point Point, matrix [][]rune) bool {
	return point.x >= 0 && point.x < len(matrix) && point.y >= 0 && point.y < len(matrix[0]) && matrix[point.x][point.y] == '.'
}

func moveRobots(matrix [][]rune, upLeftPoint Point, upLeftSymbol rune, downRightPoint Point, downRightSymbol rune) [][]rune {
	matrix = moveOneRobot(matrix, upLeftPoint, upLeftSymbol, -1, Point{0, 0})
	matrix = moveOneRobot(matrix, downRightPoint, downRightSymbol, 1, Point{len(matrix) - 1, len(matrix[0]) - 1})

	return matrix
}

func moveOneRobot(matrix [][]rune, point Point, symbol rune, firstDirection int, endPoint Point) [][]rune {
	currentPoint := point
	for {
		if currentPoint.x == endPoint.x && currentPoint.y == endPoint.y {
			break
		}

		newPoint := Point{currentPoint.x + firstDirection, currentPoint.y}
		if isCanMove(newPoint, matrix) {
			matrix[newPoint.x][newPoint.y] = symbol
			currentPoint.x = newPoint.x
			continue
		}

		newPoint = Point{currentPoint.x, currentPoint.y + firstDirection}
		if isCanMove(newPoint, matrix) {
			matrix[newPoint.x][newPoint.y] = symbol
			currentPoint.y = newPoint.y
			continue
		}
	}

	return matrix
}
