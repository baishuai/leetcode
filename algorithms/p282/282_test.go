package p282

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []string{"1+2+3", "1*2*3"}, addOperators("123", 6))
}

func Test1(t *testing.T) {
	assert.Equal(t, []string{"2+3*2", "2*3+2"}, addOperators("232", 8))
}

func Test2(t *testing.T) {
	assert.Equal(t, []string{"1*0+5", "10-5"}, addOperators("105", 5))
}

func Test3(t *testing.T) {
	assert.Equal(t, []string{}, addOperators("3456237490", 9191))
}

func Test4(t *testing.T) {
	assert.Equal(t, []string{"0+0", "0-0", "0*0"}, addOperators("00", 0))
}
