package p071

import (
	"bytes"
)

/**
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
 */

func simplifyPath(path string) string {
	stack := make([]string, 0)
	bpath := []byte(path)
	var buf bytes.Buffer
	for i := 0; i <= len(bpath); i++ {
		if i == len(bpath) || bpath[i] == '/' {
			dir := buf.String()
			buf.Reset()
			if dir == "." || len(dir) == 0 {
				continue
			} else if dir == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, dir)
			}
		} else {
			buf.WriteByte(bpath[i])
		}
	}
	buf.Reset()
	if len(stack) == 0 {
		stack = append(stack, "")
	}
	for _, dir := range stack {
		buf.WriteByte('/')
		buf.WriteString(dir)
	}
	return buf.String()
}
