package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main_old() {

	//start := time.Now()
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countNumbers int
	//fmt.Fscan(in, &countNumbers)
	fmt.Fscanf(in, "%d\n", &countNumbers)

	//fmt.Fprintln(out, "====")
	//inputNumbers3(countNumbers)
	inputNumbers2(countNumbers)

	//duration := time.Since(start)
	//fmt.Printf("Время выполнения: %d мс\n", duration.Milliseconds())
	//fmt.Fprintln(out, inputNumbers+inputNumbers)
}

func inputNumbers2(countNumbers int) {
	numbers := make([]string, countNumbers)

	for i := 0; i < countNumbers; i++ {
		var str string
		fmt.Fscan(in, &str)
		numbers[i] = createMostBigestNumber3(str)
		fmt.Fprintln(out, numbers[i])
	}
}

func inputNumbers3(countNumbers int) {
	var in *bufio.Reader
	var out *bufio.Writer

	//numbers := make([]int, countNumbers)
	numbers := make([]big.Int, countNumbers)

	for i := 0; i < countNumbers; i++ {
		_, err := fmt.Fscan(in, &numbers[i])
		if err != nil {
			fmt.Fprintln(out, err)
			break
		}
		splitNumberIntoDigits3(&numbers[i])
		//doNothing(numberByDigits)

		//fmt.Fprintln(out, &numbers[i])
	}
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
		//digits = append([]int{int(remainder.Int64())}, digits...)
	}

	return digits
}

func splitNumberIntoDigits2(n *big.Int) []int {
	temp := new(big.Int).Set(n)

	if temp.Sign() < 0 {
		temp.Abs(temp)
	}

	var digits []int

	ten := big.NewInt(10)
	remainder := new(big.Int)

	temp.DivMod(temp, ten, remainder)

	rightOne := new(big.Int).Set(remainder)
	rightReal := new(big.Int).Set(remainder)
	multiplyer := big.NewInt(10)
	result := new(big.Int)

	fmt.Fprintln(out, ">>>>>")
	fmt.Fprintln(out, temp)
	fmt.Fprintln(out, remainder)
	fmt.Fprintln(out, "<<<<<")

	for temp.Cmp(big.NewInt(0)) > 0 {
		temp.DivMod(temp, ten, remainder)

		right := new(big.Int)
		right.Mul(remainder, multiplyer)
		rightReal.Add(rightReal, right)

		multiplyer.Mul(multiplyer, ten)

		if temp.Cmp(rightOne) == -1 || temp.Cmp(rightOne) == 0 {

			left := new(big.Int)
			left.Mul(remainder, multiplyer)
			result.Add(left, rightReal)
		}

		rightOne = new(big.Int).Set(temp)

		digits = append([]int{int(remainder.Int64())}, digits...)
	}

	fmt.Fprintln(out, "res >>>>>")
	fmt.Fprintln(out, temp)
	fmt.Fprintln(out, remainder)
	fmt.Fprintln(out, "<<<<< res")

	return digits
}

func splitNumberIntoDigits3(n *big.Int) {
	temp := new(big.Int).Set(n)

	if temp.Sign() < 0 {
		temp.Abs(temp)
	}

	ten := big.NewInt(10)
	remainder := new(big.Int)

	rightReal := big.NewInt(0)

	result := big.NewInt(0)
	multiplyerLeft := big.NewInt(1)

	for temp.Cmp(big.NewInt(0)) > 0 {
		temp.DivMod(temp, ten, remainder)

		left := new(big.Int).Mul(temp, multiplyerLeft)
		ring := new(big.Int).Set(rightReal)

		resultTemp := new(big.Int).Add(left, ring)

		if resultTemp.Cmp(result) == 1 {
			result = new(big.Int).Set(resultTemp)
		}

		rightReal.Add(rightReal, new(big.Int).Mul(remainder, multiplyerLeft))
		multiplyerLeft.Mul(multiplyerLeft, ten)
	}

	fmt.Fprintln(out, result)
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

func createMostBigestNumber(numberString string) string {
	if len(numberString) <= 1 {
		return "0"
	}
	minIndex := len(numberString) - 1

	for i := 0; i < len(numberString)-1; i++ {
		if numberString[i] < numberString[i+1] {
			minIndex = i
			break
		}
	}

	result3 := numberString[0:minIndex] + numberString[minIndex+1:]

	return result3
}

func createMostBigestNumber3(numberString string) string {
	lenString := len(numberString)
	if lenString <= 1 {
		return "0"
	}

	minIndex := lenString - 1
	for i := 1; i < lenString-1; i += 2 {
		if numberString[i-1] < numberString[i] {
			minIndex = i - 1
			break
		} else if i+1 < lenString && numberString[i] < numberString[i+1] {
			minIndex = i
			break
		}
	}

	return numberString[0:minIndex] + numberString[minIndex+1:]
}
