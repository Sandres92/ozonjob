package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Arrivals struct {
	index   int
	arrival int
}

type Truck struct {
	index    int
	start    int
	end      int
	capacity int
}

type Element struct {
	Value int // Значение элемента
	Index int // Изначальный индекс элемента
}

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countDataSets4 int
	fmt.Fscanf(in, "%d\n", &countDataSets4)

	//inputDataSets4(countDataSets4)
	//result := inputDataSets4_2(countDataSets4)
	tsetOne()

	//for _, arr := range result {
	//	for i := 0; i < len(arr)-1; i++ {
	//		fmt.Fprintf(out, "%d ", arr[i])
	//	}
	//	fmt.Fprintln(out, arr[len(arr)-1])
	//}

	//testInput()
	//testFor4()

	//testRead(countDataSets4)
}

func tsetOne() {
	arrivals, tracks := inputOneDataSets4_struct()

	resultsAll := calculatePlanner(arrivals, tracks)

	for i := 0; i < len(resultsAll)-1; i++ {
		fmt.Fprintf(out, "%d ", resultsAll[i])
	}
	fmt.Fprintln(out, resultsAll[len(resultsAll)-1])
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
		fmt.Fprintln(out, "new dataset ")

		//arrivals, tracks, arrivalsIndexes, truckIndexes := inputOneDataSets4()

		//r _, track := range tracks {
		//	fmt.Fprintf(out, "%d ", track)
		//	fmt.Fprintln(out, "end ")
		//

		//resultsAll[i] = calculatePlanner(arrivals, arrivalsIndexes, tracks, truckIndexes)
	}

	return resultsAll
}

func inputOneDataSets4_struct() ([]Arrivals, []Truck) {
	var arrivalsNumbers int
	_, err := fmt.Fscanf(in, "%d\n", &arrivalsNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	//fmt.Fprintln(out, arrivalsNumbers)

	arrivals := make([]Arrivals, arrivalsNumbers)
	arrivalsIndexes := make([]int, arrivalsNumbers)
	for i := 0; i < arrivalsNumbers; i++ {
		fmt.Fscanf(in, "%d", &arrivals[i].arrival)
		arrivalsIndexes[i] = i
		arrivals[i].index = i
	}

	in.ReadString('\n') // Очищаем остатки строки после массива

	var trucksNumbers int
	_, err = fmt.Fscanf(in, "%d", &trucksNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	//fmt.Fprintf(out, "trucksNumbers = %d \n", trucksNumbers)
	in.ReadString('\n')

	trucks := make([]Truck, trucksNumbers)

	for i := 0; i < trucksNumbers; i++ {
		truckOne := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscanf(in, "%d", &truckOne[j])
		}
		in.ReadString('\n')

		trucks[i].index = i
		trucks[i].start = truckOne[0]
		trucks[i].end = truckOne[1]
		trucks[i].capacity = truckOne[2]
	}

	return arrivals, trucks
}

func inputOneDataSets4() ([]int, [][]int, []int, []int) {
	var arrivalsNumbers int
	_, err := fmt.Fscanf(in, "%d\n", &arrivalsNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	//fmt.Fprintln(out, arrivalsNumbers)

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

	//fmt.Fprintf(out, "trucksNumbers = %d \n", trucksNumbers)
	in.ReadString('\n')

	trucks := make([][]int, 3)

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

func calculatePlanner(arrivals []Arrivals, trucks []Truck) []int {
	resultPlanner := make([]int, len(arrivals))

	sort.Slice(arrivals, func(i, j int) bool {
		if arrivals[i].arrival == arrivals[j].arrival {
			return arrivals[i].index < arrivals[j].index // Сравниваем по индексу
		}
		return arrivals[i].arrival < arrivals[j].arrival // Сравниваем по значению
	})

	sort.Slice(trucks, func(i, j int) bool {
		if trucks[i].start == trucks[j].start {
			return trucks[i].index < trucks[j].index // Сравниваем по индексу
		}
		return trucks[i].start < trucks[j].start // Сравниваем по значению
	})

	//fmt.Fprintln(out, arrivals)
	//fmt.Fprintln(out)
	//fmt.Fprintln(out, trucks)

	for i := 0; i < len(arrivals); i++ {
		arrivalsIndex := arrivals[i].index
		resultPlanner[arrivalsIndex] = -1

		for j := 0; j < len(trucks); j++ {
			if arrivals[i].arrival >= trucks[j].start && arrivals[i].arrival <= trucks[j].end && trucks[j].capacity > 0 {
				resultPlanner[arrivalsIndex] = trucks[j].index + 1
				trucks[j].capacity--

				if trucks[j].capacity == 0 {
					trucks = append(trucks[:j], trucks[j+1:]...)
				}
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

func mergeSort(arr []int, arrInd []int) ([]int, []int) {
	if len(arr) <= 1 {
		return arr, arrInd
	}

	mid := len(arr) / 2
	left, leftInd := mergeSort(arr[:mid], arrInd[:mid])
	right, rightInd := mergeSort(arr[mid:], arrInd[mid:])

	return merge(left, right, leftInd, rightInd)
}

func merge(left, right []int, leftInd, rightInd []int) ([]int, []int) {
	result := []int{}
	resultInd := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			resultInd = append(resultInd, leftInd[i])
			i++
		} else {
			result = append(result, right[j])
			resultInd = append(resultInd, rightInd[i])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	resultInd = append(resultInd, leftInd[i:]...)
	resultInd = append(resultInd, rightInd[j:]...)

	return result, resultInd
}

func sortWithIndices(arr []int) ([]int, []int) {
	// Создаём срез структур Element
	elements := make([]Element, len(arr))
	for i, v := range arr {
		elements[i] = Element{Value: v, Index: i}
	}

	// Сортируем элементы
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].Value == elements[j].Value {
			return elements[i].Index < elements[j].Index // Сравниваем по индексу
		}
		return elements[i].Value < elements[j].Value // Сравниваем по значению
	})

	// Извлекаем отсортированные значения и индексы
	sortedValues := make([]int, len(arr))
	sortedIndices := make([]int, len(arr))
	for i, e := range elements {
		sortedValues[i] = e.Value
		sortedIndices[i] = e.Index
	}

	return sortedValues, sortedIndices
}
