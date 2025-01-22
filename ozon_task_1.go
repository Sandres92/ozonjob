package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {

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
