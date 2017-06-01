package p221

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	matrix := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	assert.Equal(t, 4, maximalSquare(matrix))
}
