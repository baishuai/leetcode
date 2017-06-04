package p218

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	keyPoints := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}
	assert.Equal(t, keyPoints, getSkyline(buildings))
}

func Test1(t *testing.T) {
	buildings := [][]int{}
	assert.Equal(t, [][]int{}, getSkyline(buildings))
}
