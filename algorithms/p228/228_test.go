package p228

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.Equal(t, []string{}, summaryRanges(nil))
	assert.Equal(t, []string{}, summaryRanges([]int{}))
}

func Test1(t *testing.T) {
	assert.Equal(t, []string{"1"}, summaryRanges([]int{1}))
	assert.Equal(t, []string{"1->2", "4"}, summaryRanges([]int{1, 2, 4}))
}

func Test2(t *testing.T) {
	assert.Equal(t, []string{"1->3", "5->7"}, summaryRanges([]int{1, 2, 3, 5, 6, 7}))
}

func Test3(t *testing.T) {
	assert.Equal(t, []string{"1", "3"}, summaryRanges([]int{1, 3}))
}
