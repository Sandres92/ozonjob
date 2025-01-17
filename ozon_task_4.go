package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

type Truck struct {
	start    int
	end      int
	capacity int
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countDataSets4 int
	fmt.Fscanf(in, "%d\n", &countDataSets4)

	//inputDataSets4(countDataSets4)
	result := inputDataSets4_2(countDataSets4)

	for _, arr := range result {
		for i := 0; i < len(arr); i++ {

		}
	}
	//testInput()
	//testFor4()

	//testRead(countDataSets4)
}

func testRead(countDataSets4 int) {
	a := make([]int, countDataSets4)
	for i := 0; i < countDataSets4; i++ {
		fmt.Fscan(in, &a[i])
	}

	fmt.Fprint(out, a)
}

func inputDataSets4(countDataSets4 int) []int {
	var viruses []int = make([]int, countDataSets4)

	for i := 0; i < countDataSets4; i++ {
		lineStr_1, _ := in.ReadString('\n')

		arrivalsCount, _ := strconv.Atoi(lineStr_1)
		fmt.Fprintln(out, arrivalsCount)

		lineStr_2, _ := in.ReadString('\n')
		strs := strings.Split(lineStr_2, " ")
		arrivals := make([]int, len(strs))

		for j := 0; j < len(arrivals); j++ {
			num, _ := strconv.Atoi(strs[j])
			arrivals[j] = num
		}

		lineStr_3, _ := in.ReadString('\n')
		trucksCount, _ := strconv.Atoi(lineStr_3)

		for j := 0; j < trucksCount; j++ {
			num, _ := strconv.Atoi(strs[j])
			arrivals[j] = num
		}
	}

	return viruses
}

func inputDataSets4_2(countDataSets4 int) [][]int {
	var resultsAll [][]int = make([][]int, countDataSets4)

	for i := 0; i < countDataSets4; i++ {
		arrivals, tracks, arrivalsIndexes, truckIndexes := inputOneDataSets4()
		resultsAll[i] = calculatePlanner(arrivals, arrivalsIndexes, tracks, truckIndexes)
	}

	return resultsAll
}

func inputOneDataSets4() ([]int, [][]int, []int, []int) {
	var arrivalsNumbers int
	_, err := fmt.Fscanf(in, "%d\n", &arrivalsNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	fmt.Fprintln(out, arrivalsNumbers)

	arrivals := make([]int, arrivalsNumbers)
	arrivalsIndexes := make([]int, arrivalsNumbers)
	for i := 0; i < arrivalsNumbers; i++ {
		fmt.Fscanf(in, "%d", &arrivals[i])
		arrivalsIndexes[i] = i
	}

	in.ReadString('\n') // Очищаем остатки строки после массива

	var trucksNumbers int
	_, err = fmt.Fscanf(in, "%d", &trucksNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	fmt.Fprintf(out, "trucksNumbers = %d \n", trucksNumbers)
	in.ReadString('\n')

	trucks := make([][]int, trucksNumbers)

	trucksStart := make([]int, trucksNumbers)
	trucksEnd := make([]int, trucksNumbers)
	trucksCapacity := make([]int, trucksNumbers)
	truckIndexes := make([]int, trucksNumbers)
	for i := 0; i < trucksNumbers; i++ {
		truckOne := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscanf(in, "%d", &truckOne[j])
		}
		in.ReadString('\n')

		trucksStart[i] = truckOne[0]
		trucksEnd[i] = truckOne[1]
		trucksCapacity[i] = truckOne[2]
		truckIndexes[i] = i
	}

	trucks[0] = trucksStart
	trucks[1] = trucksEnd
	trucks[2] = trucksCapacity

	return arrivals, trucks, arrivalsIndexes, truckIndexes
}

func testInput() {
	var size int
	fmt.Fscanf(os.Stdin, "%d\n", &size) // Считываем размер массива

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Fscanf(os.Stdin, "%d", &numbers[i]) // Считываем элементы массива
	}

	in.ReadString('\n') // Очищаем остатки строки после массива
	// Считываем еще одно число
	var nextNumber int
	fmt.Fscanf(os.Stdin, "%d", &nextNumber)

	// Вывод результата
	fmt.Println("Размер массива:", size)
	fmt.Println("Массив:", numbers)
	fmt.Println("Следующее число:", nextNumber)
}

func testFor4() {
	arrivals := []int{1, 9, 2, 6, 4}
	arrivalsIndexes := []int{0, 1, 2, 3, 4}
	//trucks := []Truck{{1, 8, 3}, {3, 10, 1}, {4, 7, 4}}
	arrivalsSort := make([]int, len(arrivals))
	copy(arrivalsSort, arrivals)
	arrivalsSort, arrivalsIndexes = qSort(arrivalsSort, 0, len(arrivalsSort)-1, arrivalsIndexes)

	trucksStart := []int{1, 3, 4}
	trucksEnd := []int{8, 10, 7}
	trucksCapacity := []int{3, 1, 4}
	truckIndexes := []int{0, 1, 2}
	trucksStartSort := make([]int, len(trucksStart))
	copy(trucksStartSort, trucksStart)
	trucksStartSort, truckIndexes = qSort(trucksStartSort, 0, len(trucksStart)-1, truckIndexes)

	fmt.Fprintln(out, arrivals)
	fmt.Fprintln(out, arrivalsSort)
	fmt.Fprintln(out, arrivalsIndexes)
	fmt.Fprintln(out)
	fmt.Fprintln(out, trucksStart)
	fmt.Fprintln(out, trucksStartSort)
	fmt.Fprintln(out, truckIndexes)

	resultIndexes := make([]int, len(arrivals))
	for i := 0; i < len(arrivalsIndexes); i++ {
		arrivalsIndex := arrivalsIndexes[i]
		resultIndexes[arrivalsIndex] = -1

		for j := 0; j < len(truckIndexes); j++ {
			truckIndex := truckIndexes[j]

			if arrivals[arrivalsIndex] >= trucksStart[truckIndex] && arrivals[arrivalsIndex] <= trucksEnd[truckIndex] && trucksCapacity[truckIndex] > 0 {
				resultIndexes[arrivalsIndex] = truckIndex + 1
				trucksCapacity[truckIndex]--
				break
			}
		}
	}

	for i := 0; i < len(resultIndexes); i++ {
		fmt.Fprintln(out, resultIndexes[i])
	}
}

func calculatePlanner(arrivals []int, arrivalsIndexes []int, trucks [][]int, truckIndexes []int) []int {
	resultPlanner := make([]int, len(arrivals))

	//trucks := []Truck{{1, 8, 3}, {3, 10, 1}, {4, 7, 4}}
	arrivalsSort := make([]int, len(arrivals))
	copy(arrivalsSort, arrivals)
	arrivalsSort, arrivalsIndexes = qSort(arrivalsSort, 0, len(arrivalsSort)-1, arrivalsIndexes)

	trucksStart := trucks[0]
	trucksEnd := trucks[1]
	trucksCapacity := trucks[2]

	trucksStartSort := make([]int, len(trucksStart))
	copy(trucksStartSort, trucksStart)
	trucksStartSort, truckIndexes = qSort(trucksStartSort, 0, len(trucksStart)-1, truckIndexes)

	for i := 0; i < len(arrivalsIndexes); i++ {
		arrivalsIndex := arrivalsIndexes[i]
		resultPlanner[arrivalsIndex] = -1

		for j := 0; j < len(truckIndexes); j++ {
			truckIndex := truckIndexes[j]

			if arrivals[arrivalsIndex] >= trucksStart[truckIndex] && arrivals[arrivalsIndex] <= trucksEnd[truckIndex] && trucksCapacity[truckIndex] > 0 {
				resultPlanner[arrivalsIndex] = truckIndex + 1
				trucksCapacity[truckIndex]--
				break
			}
		}
	}

	return resultPlanner
}

func qSort(lst []int, left int, right int, arrInd []int) ([]int, []int) {
	l := left
	r := right
	center := lst[(left+right)/2]
	for l <= r {
		for lst[r] > center {
			r--
		}
		for lst[l] < center {
			l++
		}
		if l <= r {
			lst[r], lst[l] = lst[l], lst[r]
			arrInd[r], arrInd[l] = arrInd[l], arrInd[r]
			l++
			r--
		}
	}
	if r > left {
		qSort(lst, left, r, arrInd)
	}
	if l < right {
		qSort(lst, l, right, arrInd)
	}
	return lst, arrInd
}
