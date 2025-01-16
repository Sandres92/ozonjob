package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main3() {

	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	fmt.Fscan(in, &str)

	//numbers := strings.Split(str, "\n")

	//lenStr := len(numbers)

	//fmt.Fprint(out, numbers)
	//fmt.Fprint(out, "ddd = %d", lenStr)

	//str := "strawberry\n blueberry\n raspberry"
	fmt.Printf("strings.Split(): %#v\n", strings.Split(str, "\n"))

}
