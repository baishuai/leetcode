package p434

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	assert.Equal(t, 0, countSegments(""))
}

func TestE1(t *testing.T) {
	assert.Equal(t, 6, countSegments(", , , ,        a, eaefa"))
}
