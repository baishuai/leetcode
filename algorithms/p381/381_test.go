package p381

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {

	rc := Constructor()
	assert.Equal(t, true, rc.Insert(0))
	assert.Equal(t, true, rc.Remove(0))
	assert.Equal(t, true, rc.Insert(-1))
	assert.Equal(t, false, rc.Remove(0))
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
	assert.Equal(t, -1, rc.GetRandom())
}

func Test1(t *testing.T) {
	//["RandomizedCollection","insert","insert","remove","insert","remove","getRandom"]
	//[[],[0],[1],[0],[2],[1],[]]

	rc := Constructor()
	assert.Equal(t, true, rc.Insert(0))
	assert.Equal(t, true, rc.Insert(1))
	assert.Equal(t, true, rc.Remove(0))
	assert.Equal(t, true, rc.Insert(2))
	assert.Equal(t, true, rc.Remove(1))
	assert.Equal(t, 2, rc.GetRandom())
	assert.Equal(t, 1, len(rc.vec))
}

func Test2(t *testing.T) {
	//["RandomizedCollection","insert","insert", "insert","insert","insert", "remove","remove", "remove","insert","remove","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom"]
	//[[],[1],[1], [2],[2],[2], [1],[1], [2],[1],[2],[],[],[],[],[],[],[],[],[],[]]
	rc := Constructor()
	assert.Equal(t, true, rc.Insert(1))
	assert.Equal(t, false, rc.Insert(1))

	assert.Equal(t, true, rc.Insert(2))
	assert.Equal(t, false, rc.Insert(2))
	assert.Equal(t, false, rc.Insert(2))

	assert.Equal(t, true, rc.Remove(1))
	assert.Equal(t, true, rc.Remove(1))

	assert.Equal(t, true, rc.Remove(2))
	assert.Equal(t, true, rc.Insert(1))

	t.Log(rc.vec)
	assert.Equal(t, true, rc.Remove(2))

	assert.Equal(t, 2, len(rc.vec))
	rc.GetRandom()
	rc.GetRandom()

}
