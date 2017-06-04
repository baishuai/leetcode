package p295

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	assert.Equal(t, 1.5, m.FindMedian()) //-> 1.5
	m.AddNum(3)
	assert.Equal(t, 2.0, m.FindMedian()) //- > 2
}

func Test1(t *testing.T) {
	m := Constructor()
	m.AddNum(-1)
	assert.Equal(t, -1.0, m.FindMedian())
	m.AddNum(-2)
	assert.Equal(t, -1.5, m.FindMedian())
	m.AddNum(-3)
	assert.Equal(t, -2.0, m.FindMedian())
	m.AddNum(-4)
	assert.Equal(t, -2.5, m.FindMedian())
	m.AddNum(-5)
	assert.Equal(t, -3.0, m.FindMedian())
}

func Test2(t *testing.T) {
	m := Constructor()
	m.AddNum(12)
	//assert.Equal(t, -1.0, m.FindMedian())
	m.AddNum(10)
	//assert.Equal(t, -1.5, m.FindMedian())
	m.AddNum(13)
	//assert.Equal(t, -2.0, m.FindMedian())
	m.AddNum(11)
	//assert.Equal(t, -2.5, m.FindMedian())
	m.AddNum(5)
	//assert.Equal(t, -3.0, m.FindMedian())
}
