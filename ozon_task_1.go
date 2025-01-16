package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	//"time"
)

var in *bufio.Reader
var out *bufio.Writer

const base = 10.0
const exponent = 5.0

func main() {
	//start := time.Now()
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countNumbers int
	fmt.Fscan(in, &countNumbers)

	if countNumbers < 1 || countNumbers > int(math.Pow(base, exponent)) {
		fmt.Fprintln(out, "Error input, number less 0 or more 10^5")
		return
	}

	//fmt.Fprintln(out, "====")
	numbers := inputNumbers(countNumbers)

	if numbers == nil {
		fmt.Fprintln(out, "Error input numbers")
	} else {
		for i := 0; i < countNumbers; i++ {
			fmt.Fprintln(out, &numbers[i])
		}
	}

	//duration := time.Since(start)
	//fmt.Printf("Время выполнения: %d мс\n", duration.Milliseconds())
	//fmt.Fprintln(out, inputNumbers+inputNumbers)
}

func inputNumbers(countNumbers int) []big.Int {
	//numbers := make([]int, countNumbers)
	numbers := make([]big.Int, countNumbers)

	for i := 0; i < countNumbers; i++ {
		_, err := fmt.Fscan(in, &numbers[i])
		if err != nil {
			fmt.Fprintln(out, err)
			return nil
		}
		numberByDigits := splitNumberIntoDigits(&numbers[i])

		if len(numberByDigits) <= 0 {
			continue
		}
		if len(numberByDigits) > int(math.Pow(base, exponent)) {
			continue
		}
		//if len(numberByDigits) <= 0 || len(numberByDigits) > int(math.Pow(base, exponent)) {
		//	fmt.Fprintln(out, "Numbel size less 0 or more 10^5")
		//	continue
		//}
		numberByDigits = removeLessDigit(numberByDigits)

		numbers[i] = *assembleNumberFromDigits(numberByDigits)
	}
	return numbers
}

func splitNumberIntoDigits(n *big.Int) []int {
	temp := new(big.Int).Set(n)

	if temp.Sign() < 0 {
		temp.Abs(temp)
	}

	var digits []int

	ten := big.NewInt(10)
	remainder := new(big.Int)

	for temp.Cmp(big.NewInt(0)) > 0 {
		temp.DivMod(temp, ten, remainder)
		digits = append([]int{int(remainder.Int64())}, digits...)
	}

	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	return digits
}

func assembleNumberFromDigits(digits []int) *big.Int {
	result := big.NewInt(0)
	ten := big.NewInt(10)

	for _, digit := range digits {
		result.Mul(result, ten)
		result.Add(result, big.NewInt(int64(digit)))
	}

	return result
}

func removeLessDigit(numberByDigits []int) []int {
	if len(numberByDigits) == 0 {
		arr := []int{0}
		return arr
	}

	minIndex := 0
	for i := 1; i < len(numberByDigits); i++ {
		if numberByDigits[i] < numberByDigits[minIndex] {
			minIndex = i
		}
	}

	result := append(numberByDigits[:minIndex], numberByDigits[minIndex+1:]...)
	return result
}
