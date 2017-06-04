package p552

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, checkRecord(1))
	assert.Equal(t, 8, checkRecord(2))
}

func Test50(t *testing.T) {
	assert.Equal(t, 100469819, checkRecord(50))

}

func Test100(t *testing.T) {
	assert.Equal(t, 985598218, checkRecord(100))
}
