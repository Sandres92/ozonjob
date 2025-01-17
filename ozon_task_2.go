package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		secondString = strings.TrimRight(secondString, "\r\n")

		_, err := strconv.Atoi(zeroString)

		if err != nil {
			results[i] = "no"
		} else if results[i] == "yes" {
			re := regexp.MustCompile("[0-9]+$")

			if len(firstString) != len(secondString) {
				results[i] = "no"
			} else if !re.MatchString(firstString) || !re.MatchString(secondString) {
				results[i] = "no"
			} else {
				resultsStr := workOutString(firstString)

				if resultsStr != secondString {
					results[i] = "no"
				}
			}
		}
	}

	return results
}

func workOutString(inputStr string) string {
	strs := strings.Split(inputStr, " ")
	numbers := make([]int, len(strs))

	for i := 0; i < len(numbers); i++ {
		num, _ := strconv.Atoi(strs[i])
		numbers[i] = num
	}

	sort.Ints(numbers)

	var resultStr strings.Builder
	for i := 0; i < len(numbers)-1; i++ {
		str := strconv.Itoa(numbers[i])

		resultStr.WriteString(str)
		resultStr.WriteString(" ")
	}

	resultStr.WriteString(strconv.Itoa(numbers[len(numbers)-1]))

	return resultStr.String()
}
