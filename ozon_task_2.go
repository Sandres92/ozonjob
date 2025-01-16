package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	//duration := time.Since(start)
	//fmt.Printf("Время выполнения: %d мс\n", duration.Milliseconds())
	//fmt.Fprintln(out, inputNumbers+inputNumbers)
}
