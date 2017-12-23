package p726

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test0(t *testing.T) {
// 	formula := "H2O"
// 	exp := "H2O"
// 	assert.Equal(t, exp, countOfAtoms(formula))
// }

func Test1(t *testing.T) {
	formula := "K4(ON(SO3)2)2"
	exp := "K4N2O14S4"
	assert.Equal(t, exp, countOfAtoms(formula))
}
