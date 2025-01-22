package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
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

func main4() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countDataSets4 int
	fmt.Fscanf(in, "%d\n", &countDataSets4)
	start := time.Now()
	//inputDataSets4(countDataSets4)
	result := inputDataSets(countDataSets4)
	//tsetOne()
	duration := time.Since(start)
	fmt.Fprintf(out, "all input and calc = %d\n", duration.Milliseconds())

	start = time.Now()
	for _, arr := range result {
		for i := 0; i < len(arr)-1; i++ {
			fmt.Fprintf(out, "%d ", arr[i])
		}
		fmt.Fprintln(out, arr[len(arr)-1])
	}

	duration = time.Since(start)
	fmt.Fprintf(out, "output = %d\n", duration.Milliseconds())
	//testInput()
	//testFor4()

	//testRead(countDataSets4)
}

func tsetOne() {
	arrivals, tracks := inputOneDataSets()

	resultsAll := calculatePlanner(arrivals, tracks)

	for i := 0; i < len(resultsAll)-1; i++ {
		fmt.Fprintf(out, "%d ", resultsAll[i])
	}
	fmt.Fprintln(out, resultsAll[len(resultsAll)-1])
}

func inputDataSets(countDataSets4 int) [][]int {
	var resultsAll [][]int = make([][]int, countDataSets4)

	for i := 0; i < countDataSets4; i++ {
		start := time.Now()

		arrivals, tracks := inputOneDataSets()
		duration := time.Since(start)
		fmt.Fprintf(out, "%d input = %d\n", i, duration.Nanoseconds())

		start = time.Now()
		resultsAll[i] = calculatePlanner(arrivals, tracks)
		duration = time.Since(start)
		fmt.Fprintf(out, "%d calc = %d\n", i, duration.Nanoseconds())
	}

	return resultsAll
}

func inputOneDataSets() ([]Arrivals, []Truck) {
	var arrivalsNumbers int
	_, err := fmt.Fscanf(in, "%d\n", &arrivalsNumbers)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	//fmt.Fprintln(out, arrivalsNumbers)

	arrivals := make([]Arrivals, arrivalsNumbers)
	//arrivalsIndexes := make([]int, arrivalsNumbers)
	for i := 0; i < arrivalsNumbers; i++ {
		arrivals[i].index = i
		fmt.Fscanf(in, "%d", &arrivals[i].arrival)
		//arrivalsIndexes[i] = i
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
		//truckOne := make([]int, 3)
		//for j := 0; j < 3; j++ {
		//}

		trucks[i].index = i

		fmt.Fscanf(in, "%d %d %d", &trucks[i].start, &trucks[i].end, &trucks[i].capacity)
		in.ReadString('\n')

		//trucks[i].start = truckOne[0]
		//trucks[i].end = truckOne[1]
		//trucks[i].capacity = truckOne[2]
	}

	return arrivals, trucks
}

func calculatePlanner(arrivals []Arrivals, trucks []Truck) []int {

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

	//trackEndCapacity := len(trucks)

	resultPlanner := make([]int, len(arrivals))
	currcntArrivalIndex := 0

	for j := 0; j < len(trucks); j++ {

		for i := currcntArrivalIndex; i < len(arrivals); i++ {

			if arrivals[i].arrival < trucks[j].start {
				currcntArrivalIndex++
				resultPlanner[arrivals[i].index] = -1

				continue
			}

			if arrivals[i].arrival >= trucks[j].start && arrivals[i].arrival <= trucks[j].end && trucks[j].capacity > 0 {
				resultPlanner[arrivals[i].index] = trucks[j].index + 1
				trucks[j].capacity--
				currcntArrivalIndex++

				if trucks[j].capacity == 0 {
					break
				}
			} else {
				//resultPlanner[arrivals[i].index] = -1
				break
			}
		}
	}

	//fmt.Fprintf(out, "currcntArrivalIndex  %d \n", currcntArrivalIndex)
	for i := currcntArrivalIndex; i < len(arrivals); i++ {
		resultPlanner[arrivals[i].index] = -1
	}

	//for i := 0; i < len(arrivals); i++ {
	//	arrivalsIndex := arrivals[i].index
	//	resultPlanner[arrivalsIndex] = -1
	//
	//	for j := 0; j < len(trucks); j++ {
	//		if trucks[j].capacity < 0 {
	//			continue
	//		}
	//
	//		if arrivals[i].arrival >= trucks[j].start && arrivals[i].arrival <= trucks[j].end && trucks[j].capacity > 0 {
	//			resultPlanner[arrivalsIndex] = trucks[j].index + 1
	//			trucks[j].capacity--
	//
	//			//if trucks[j].capacity == 0 {
	//			//	trucks = append(trucks[:j], trucks[j+1:]...)
	//			//}
	//			break
	//		}
	//	}
	//}

	return resultPlanner
}
