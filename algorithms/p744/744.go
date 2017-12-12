package p744

import (
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})
	idx = idx % len(letters)
	return letters[idx]
}
