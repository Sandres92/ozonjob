package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countNumbers int

	fmt.Fscan(in, &countNumbers)
	numbers := make([]big.Int, countNumbers)

	for i := 0; i < countNumbers; i++ {
		_, err := fmt.Fscan(in, &numbers[i])
		if err != nil {
			fmt.Fprintln(out, err)
			break
		}
	}

	ten := big.NewInt(10)
	for i := 0; i < countNumbers; i++ {
		temp := new(big.Int).Set(&numbers[i])
		temp.Div(temp, ten)
		fmt.Fprintln(out, temp)
	}

	//fmt.Fprintln(out, numbers)
}
