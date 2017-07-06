package p567

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, checkInclusion("ab", "eidbaooo"))
}

func Test1(t *testing.T) {
	assert.Equal(t, false, checkInclusion("ab", "eidboaoo"))
}
