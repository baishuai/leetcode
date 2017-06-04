package p468

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, "Neither", validIPAddress("1081:db8:85a3:01:-0:8A2E:0370:7334"))
}
