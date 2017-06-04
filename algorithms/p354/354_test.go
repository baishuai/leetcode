package p354

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	envs := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	assert.Equal(t, 3, maxEnvelopes(envs))
}

func Test1(t *testing.T) {
	envs := [][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}
	assert.Equal(t, 5, maxEnvelopes(envs))
}
