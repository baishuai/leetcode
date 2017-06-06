package p551

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, checkRecord("PPALLP"))
	assert.Equal(t, false, checkRecord("PPALLPA"))
	assert.Equal(t, false, checkRecord("PPALLL"))
}
