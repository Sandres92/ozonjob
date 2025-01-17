package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {
	//start := time.Now()
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//fmt.Fscan(in, &countNumbers)

	number := int(math.Pow(10.0, 3.0))
	fmt.Fprintln(out, number)

	var countNumbers int
	fmt.Fscanf(in, "%d\n", &countNumbers)

	resultsArr := inputDataSet(countNumbers)

	for i := 0; i < len(resultsArr); i++ {
		fmt.Fprintln(out, resultsArr[i])
	}
}

func inputDataSet(countDataSets int) []string {

	results := make([]string, countDataSets)

	for i := 0; i < countDataSets; i++ {
		results[i] = "yes"

		var zeroString, firstString, secondString string

		zeroString, _ = in.ReadString('\n')
		zeroString = strings.TrimSpace(zeroString) // Убираем символ новой строки

		firstString, _ = in.ReadString('\n')
		firstString = strings.TrimSpace(firstString)

		secondString, _ = in.ReadString('\n')
		secondString = strings.TrimSpace(secondString)

		_, err := strconv.Atoi(zeroString)

		if err != nil {
			fmt.Fprintln(out, "1111")
			results[i] = "no"
		} else if results[i] == "yes" {
			isDigit := func(c rune) bool { return (c >= '0' && c <= '9') || c == '_' || c == ' ' }

			fmt.Fprintln(out, "2222")
			fmt.Fprintln(out, isDigit)

			if !strings.ContainsFunc(firstString, isDigit) || !strings.ContainsFunc(secondString, isDigit) {
				results[i] = "no"
				fmt.Fprintln(out, "3333")
			}
			resultsStr := workOutString(firstString)

			if resultsStr != secondString {
				results[i] = "no"
				fmt.Fprintln(out, "4444")
			}
		}
	}

	return results
}

func workOutString(inputStr string) string {
	resultStr := ""
	strs := strings.Split(inputStr, " ")

	numbers := make([]int, len(strs))

	for i := 0; i < len(numbers); i++ {
		num, err := strconv.Atoi(strs[i])
		if err != nil {
			fmt.Fprintln(out, "error not number")
			fmt.Fprintln(out, inputStr)
			fmt.Fprintln(out, i)
			fmt.Fprintln(out, "end error")
			break
		}

		numbers[i] = num
	}

	var emptyArr []int
	copy(emptyArr, numbers)
	sort.Ints(emptyArr)

	for i := 0; i < len(emptyArr)-1; i++ {
		str := strconv.Itoa(emptyArr[i])

		resultStr += str + " "
	}

	resultStr += strconv.Itoa(emptyArr[len(emptyArr)-1])

	return resultStr
}
