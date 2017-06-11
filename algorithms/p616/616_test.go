package p616

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, "<b>abc</b>xyz<b>123</b>", addBoldTag("abcxyz123", []string{"123", "abc"}))
}
