package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	no := flag.Int("n", 0, "Problems No.")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	if _, err := os.Stat(fmt.Sprintf("p%03d", *no)); err == nil {
		// path/to/whatever exists
		fmt.Println("文件夹已存在")
		return
	}

	os.Mkdir(fmt.Sprintf("p%03d", *no), 0755)

	if f, err := os.Create(fmt.Sprintf("p%03d/%03d.go", *no, *no)); err == nil {
		f.WriteString(fmt.Sprintf("package p%03d", *no))
		f.Close()
	}
	if f, err := os.Create(fmt.Sprintf("p%03d/%03d_test.go", *no, *no)); err == nil {
		f.WriteString(fmt.Sprintf("package p%03d", *no))
		f.Close()
	}
}
