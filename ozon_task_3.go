package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

type Directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []Directory `json:"folders"`
}

func main() {
	//start := time.Now()
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//fmt.Fscan(in, &countNumbers)

	var countJsons int
	fmt.Fscanf(in, "%d\n", &countJsons)
	//testJsons(countJsons)

	resultsArr := inputJsons(countJsons)

	for i := 0; i < len(resultsArr); i++ {
		fmt.Fprintln(out, resultsArr[i])
	}
}

func inputJsons(countJsons int) []int {
	var viruses []int = make([]int, countJsons)

	for i := 0; i < countJsons; i++ {
		var countLinesJson int
		_, err := fmt.Fscanf(in, "%d\n", &countLinesJson)

		if err != nil {
			fmt.Fprintln(out, err)
			break
		}

		jsonStr := inputOneJsons(countLinesJson)

		var directory Directory
		errJson := json.Unmarshal([]byte(jsonStr), &directory)
		if errJson != nil {
			fmt.Println("Ошибка чтения JSON-данных:", err)
		}

		viruses[i] = countNumberOfViruses(directory, false)
	}

	return viruses
}

func inputOneJsons(countLinesJsons int) string {
	var jsonStr strings.Builder

	for i := 0; i < countLinesJsons; i++ {
		lineStr, _ := in.ReadString('\n')
		//lineStr = strings.TrimSpace(lineStr) // Убираем символ новой строки

		jsonStr.WriteString(lineStr)
	}

	return jsonStr.String()
}

func countNumberOfViruses(directory Directory, isVirused bool) int {
	virusCount := 0
	isVirus := isVirused

	if isVirus {
		virusCount += len(directory.Files)
	} else {
		for _, file := range directory.Files {
			ext := path.Ext(file)

			if ext == ".hack" {
				virusCount += len(directory.Files)
				isVirus = true
				break
			}
		}
	}

	for i := 0; i < len(directory.Folders); i++ {
		virusCount += countNumberOfViruses(directory.Folders[i], isVirus)
	}

	return virusCount
}

func testJsons(countJsons int) {

	var directory Directory
	jsonData := ``
	err := json.Unmarshal([]byte(jsonData), &directory)
	if err != nil {
		fmt.Println("Ошибка чтения JSON-данных:", err)
	}
	//fmt.Println(directory)

	//file := "kta.hack"
	//expand := strings.TrimLeft(file, ".")
	//fmt.Println(expand)
	//
	//ext := path.Ext(file)
	//fmt.Println(ext) // Output: ".txt"

	//listFiles(directory)
	//fmt.Println("=======")
	//virusFiles(directory, false)
	//fmt.Println(virusFilesStr)

	v := countNumberOfViruses(directory, false)
	fmt.Fprintln(out, v)
}

func listFiles(directory Directory) {
	for _, file := range directory.Files {
		fmt.Fprintln(out, file)
	}

	for i := 0; i < len(directory.Folders); i++ {
		listFiles(directory.Folders[i])
	}
}

func virusFiles(directory Directory, isVirused bool) string {
	var virusFilesStr string
	var isVirus bool

	if isVirused {
		isVirus = true
		for _, file := range directory.Files {
			virusFilesStr += file + "\n"
		}
		fmt.Printf("dir_1 %s = %d \n", directory.Dir, len(directory.Files))
	} else {

		for _, file := range directory.Files {
			ext := path.Ext(file)

			if ext == ".hack" {
				isVirus = true
				break
			}
		}

		if isVirus {
			for _, file := range directory.Files {
				virusFilesStr += file + "\n"
			}
			fmt.Printf("dir_2 %s = %d \n", directory.Dir, len(directory.Files))
		}
	}

	for i := 0; i < len(directory.Folders); i++ {
		virusFilesStr += virusFiles(directory.Folders[i], isVirus)
	}

	return virusFilesStr
}
