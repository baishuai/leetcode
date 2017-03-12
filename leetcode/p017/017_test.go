package p017

import (
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {

	ans := letterCombinations("23")
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	if ! reflect.DeepEqual(ans, expect) {
		t.Fatal("error answer", ans)

	}
}
