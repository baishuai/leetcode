package p609

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	exp := make(map[string]string)
	exp["root/a/1.txt"] = "abcd"
	exp["root/a/2.txt"] = "efgh"

	assert.Equal(t, exp, getfile("root/a 1.txt(abcd) 2.txt(efgh)"))
}

func Test1(t *testing.T) {
	exp := make(map[string][]string, 0)
	exp["root/a/1.txt"] = []string{"root/a/1.txt", "root/c/3.txt"}
	exp["root/4.txt"] = []string{"root/4.txt", "root/a/2.txt", "root/c/d/4.txt"}

	ans := findDuplicate([]string{
		"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)",
	})
	ansmap := make(map[string][]string)

	for _, paths := range ans {
		sort.Strings(paths)
		ansmap[paths[0]] = paths
	}
	assert.Equal(t, exp, ansmap)
}
