package p737

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	words1 := []string{"great", "acting", "skills"}
	words2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{
		{"great", "good"}, {"fine", "good"}, {"acting", "drama"}, {"skills", "talent"},
	}
	assert.True(t, areSentencesSimilarTwo(words1, words2, pairs))

}
