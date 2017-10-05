package p434

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.Equal(t, 0, countSegments(""));
}

func TestE1(t *testing.T) {
	assert.Equal(t,6, countSegments(", , , ,        a, eaefa"))
}