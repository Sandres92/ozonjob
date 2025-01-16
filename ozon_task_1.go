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
	//fmt.Fscan(in, &countNumbers)
	fmt.Fscanf(in, "%d\n", &countNumbers)

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
		_, err := fmt.Fscanf(in, "%d\n", &numbers[i])
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

//func inputNumbers2(countNumbers int) []big.Int {
//	//numbers := make([]int, countNumbers)
//	numbers := make([]big.Int, countNumbers)
//
//	for i := 0; i < countNumbers; i++ {
//		_, err := fmt.Fscanf(in, "%d\n", &numbers[i])
//		if err != nil {
//			fmt.Fprintln(out, err)
//			break
//		}
//	}
//
//	ten := big.NewInt(10)
//	for i := 0; i < countNumbers; i++ {
//		temp := new(big.Int).Set(&numbers[i])
//		temp.Div(temp, ten)
//
//		numbers[i] = *temp
//	}
//
//	return numbers
//}

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

	return digits
}

func assembleNumberFromDigits(digits []int) *big.Int {
	result := big.NewInt(0)
	ten := big.NewInt(10)

	for i := 0; i < len(digits); i++ {
		result.Mul(result, ten)
		result.Add(result, big.NewInt(int64(digits[i])))
	}

	return result
}

func removeLessDigit(numberByDigits []int) []int {
	if len(numberByDigits) <= 1 {
		arr := []int{0}
		return arr
	}

	minIndex := len(numberByDigits) - 1
	for i := 0; i < len(numberByDigits)-1; i++ {
		if numberByDigits[i] < numberByDigits[i+1] {
			minIndex = i
			break
		}
	}

	result := append(numberByDigits[:minIndex], numberByDigits[minIndex+1:]...)
	return result
}
