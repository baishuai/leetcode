package p299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {

	assert.Equal(t, "1A3B", getHint("1807", "7810"))
	assert.Equal(t, "1A1B", getHint("1123", "0111"))
}
