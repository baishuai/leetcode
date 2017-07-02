package main

import (
	"flag"
	"fmt"
	"os"
)

const cppmain = `
#ifndef LEETCODE_%03d_HPP
#define LEETCODE_%03d_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>

using namespace std;


#endif //LEETCODE_%03d_HPP
`

func gencppMain(id int) string {
	return fmt.Sprintf(cppmain, id, id, id)
}

const cpptest = `
#include "%03d.hpp"
#include <gtest/gtest.h>

TEST(p%03d, example) {
    Solution s;
}
`

func gencppTest(id int) string {
	return fmt.Sprintf(cpptest, id, id)
}

func main() {
	no := flag.Int("n", 0, "Problems No.")
	lang := flag.String("l", "cpp", "Language Template.")
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

	if *lang == "go" {

	} else if *lang == "cpp" {

	}

	switch *lang {
	case "go":
		if f, err := os.Create(fmt.Sprintf("p%03d/%03d.go", *no, *no)); err == nil {
			f.WriteString(fmt.Sprintf("package p%03d", *no))
			f.Close()
		}
		if f, err := os.Create(fmt.Sprintf("p%03d/%03d_test.go", *no, *no)); err == nil {
			f.WriteString(fmt.Sprintf("package p%03d", *no))
			f.Close()
		}
	case "cpp":
		if f, err := os.Create(fmt.Sprintf("p%03d/%03d.hpp", *no, *no)); err == nil {
			f.WriteString(gencppMain(*no))
			f.Close()
		}

		if f, err := os.Create(fmt.Sprintf("p%03d/%03d_test.cpp", *no, *no)); err == nil {
			f.WriteString(gencppTest(*no))
			f.Close()
		}
	default:
		fmt.Println("未知语言类型, unknown program language")
		os.RemoveAll(fmt.Sprintf("p%03d", *no))
	}

}
