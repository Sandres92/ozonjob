package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Point5_old struct {
	x, y int
}

type QueueNode struct {
	point Point5_old
	path  []Point5_old // Заранее выделяем память под маршрут
	len   int          // Длина текущего маршрута
}

func main5_old() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countMatrices int
	fmt.Fscanf(in, "%d\n", &countMatrices)

	inputMatices5_old(countMatrices)
	//inputOneMatrix()
}

func testFindPath() {

	// Пример матрицы
	matrix := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '#', 'A', '#', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '#', '.', '#', '.'},
		{'.', '.', '.', '.', '.'},
	}

	result := findPath(matrix, 5, 5, Point5_old{1, 2}, Point5_old{0, 0})
	if result != nil {
		//fmt.Fprintln(out, "Кратчайший путь имеет длину: ", len(result))

		// Преобразуем маршрут в массив координат
		fmt.Fprintf(out, "len path : %d\n", len(result.path))
		for i := 1; i < result.len; i++ {

			p := Point5_old{result.path[i].x, result.path[i].y}
			matrix[p.x][p.y] = 'a'
			fmt.Fprintf(out, "path : %d, %d\n", p.x, p.y)
		}
		fmt.Fprintln(out, "====")
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				fmt.Fprintf(out, "%s", string(matrix[i][j]))
			}
			fmt.Fprintln(out)
		}

	} else {
		fmt.Println("Маршрут недоступен.")
	}
}

func inputMatices5_old(countMatrices int) [][][]rune {
	resultMatrces := make([][][]rune, countMatrices)

	for k := 0; k < countMatrices; k++ {
		resultMatrces[k] = inputOneMatrix5_old()
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

func inputOneMatrix5_old() [][]rune {
	var sizeN, sizeM int
	_, err := fmt.Fscanf(in, "%d %d\n", &sizeN, &sizeM)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	originalMatrix := make([][]rune, sizeN)
	resultMatrix := make([][]rune, sizeN)
	var pointA, pointB Point5_old

	for i := 0; i < sizeN; i++ {

		originalMatrix[i] = make([]rune, sizeM)
		resultMatrix[i] = make([]rune, sizeM)

		for j := 0; j < sizeM; j++ {
			fmt.Fscanf(in, "%c", &originalMatrix[i][j])

			if originalMatrix[i][j] == 'A' {
				pointA.x = i
				pointA.y = j
			} else if originalMatrix[i][j] == 'B' {
				pointB.x = i
				pointB.y = j
			}
			resultMatrix[i][j] = originalMatrix[i][j]
		}
		in.ReadString('\n')
	}

	if (pointA.x == 0 && pointA.y == 0 && pointB.x == sizeN-1 && pointB.y == sizeM-1) ||
		(pointA.x == sizeN-1 && pointA.y == sizeM-1 && pointB.x == 0 && pointB.y == 0) {
		return originalMatrix
	}

	if pointA.x+pointA.y < pointB.x+pointB.y {

	} else if pointA.x+pointA.y < pointB.x+pointB.y {

	} else {
		path := findPath(originalMatrix, sizeN, sizeM, pointA, Point5_old{0, 0})
		if path != nil {
			for i := 1; i < path.len; i++ {
				p := Point5_old{path.path[i].x, path.path[i].y}
				resultMatrix[p.x][p.y] = 'a'
			}

			path = findPath(resultMatrix, sizeN, sizeM, pointB, Point5_old{sizeN - 1, sizeM - 1})

			if path != nil {
				for i := 1; i < path.len; i++ {
					p := Point5_old{path.path[i].x, path.path[i].y}
					resultMatrix[p.x][p.y] = 'b'
				}
				return resultMatrix
			}
		}

		copy(resultMatrix, originalMatrix)
		path = findPath(originalMatrix, sizeN, sizeM, pointB, Point5_old{0, 0})
		if path != nil {
			for i := 1; i < path.len; i++ {
				p := Point5_old{path.path[i].x, path.path[i].y}
				resultMatrix[p.x][p.y] = 'b'
			}

			path = findPath(resultMatrix, sizeN, sizeM, pointA, Point5_old{sizeN - 1, sizeM - 1})

			if path != nil {
				for i := 1; i < path.len; i++ {
					p := Point5_old{path.path[i].x, path.path[i].y}
					resultMatrix[p.x][p.y] = 'a'
				}
				return resultMatrix
			}
		}
	}

	return originalMatrix
}

func isValid(x, y, sizeN, sizeM int, matrix [][]rune, visited [][]bool) bool {
	return x >= 0 && x < sizeN && y >= 0 && y < sizeM && matrix[x][y] == '.' && !visited[x][y]
}

func findPath(matrix [][]rune, sizeN, sizeM int, start, end Point5_old) *QueueNode {
	initialPath := make([]Point5_old, sizeN*sizeM)
	initialPath[0] = Point5_old{start.x, start.y}

	queue := list.New()
	queue.PushBack(QueueNode{Point5_old{start.x, start.y}, initialPath, 1})

	if start.x == end.x && start.y == end.y {
		node := queue.Front().Value.(QueueNode)
		return &node
	}

	directions := []Point5_old{
		{-1, 0}, // вверх
		{1, 0},  // вниз
		{0, -1}, // влево
		{0, 1},  // вправо
	}

	visited := make([][]bool, sizeN)
	for i := range visited {
		visited[i] = make([]bool, sizeM)
	}

	visited[start.x][start.y] = true

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(QueueNode)
		current := node.point

		if current.x == end.x && current.y == end.y {
			return &node
		}

		for _, dir := range directions {
			newX := current.x + dir.x
			newY := current.y + dir.y

			if isValid(newX, newY, sizeN, sizeM, matrix, visited) {
				newPath := make([]Point5_old, len(node.path))
				copy(newPath, node.path)
				newPath[node.len] = Point5_old{newX, newY}

				queue.PushBack(QueueNode{Point5_old{newX, newY}, newPath, node.len + 1})

				visited[newX][newY] = true
			}
		}
	}

	return nil
}
