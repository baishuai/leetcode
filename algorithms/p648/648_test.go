package p648

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {

	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	exp := "the cat was rat by the bat"

	assert.Equal(t, exp, replaceWords(dict, sentence))
}
