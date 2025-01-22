package main

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_ValidateOutput(t *testing.T) {
	dir := "E:/Aleks/go/ozon/remove-digit/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var in *os.File
	var exp string

	for _, e := range entries {
		if strings.Contains(e.Name(), ".test.") {
			continue
		}

		if strings.Contains(e.Name(), ".a") {
			f, _ := os.ReadFile(dir + e.Name())
			exp = string(f)
		} else {
			in, _ = os.Open(dir + e.Name())
		}

		if in != nil && string(exp) != "" {
			t.Log(e.Name())

			oldR, oldW := os.Stdin, os.Stdout

			out2, err := os.CreateTemp(dir, e.Name()+".test.")
			if err != nil {
				t.Errorf("tmp fail %+v", err)
			}
			os.Stdin = in
			os.Stdout = out2

			start := time.Now()
			main()
			duration := time.Since(start)
			t.Logf("time = %d\n", duration.Milliseconds())
			//cmd := exec.Command("go", "run", "../ozon_task/ozon_task_1.go")

			in.Close()
			in = nil

			out2.Seek(0, io.SeekStart)
			res, _ := io.ReadAll(out2)
			out2.Close()

			os.Stdin, os.Stdout = oldR, oldW

			if exp != string(res) {
				es := strings.Split(exp, "\n")
				gs := strings.Split(string(res), "\n")

				for i := 0; i < len(es); i++ {
					if es[i] != gs[i] {
						t.Errorf("failed %s on %d\nexp: %sgot: %s", e.Name(), i, es[i], gs[i])
					}
				}
				return
			}

			os.Remove(out2.Name())
			exp = ""
		}
	}
}
