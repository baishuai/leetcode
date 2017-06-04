package p587

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Points []Point

func (sp Points) Len() int {
	return len(sp)
}

func (sp Points) Less(i, j int) bool {
	return sp[i].X < sp[j].X || sp[i].X == sp[j].X && sp[i].Y < sp[j].Y
}

func (sp Points) Swap(i, j int) {
	sp[i], sp[j] = sp[j], sp[i]
}

func Test0(t *testing.T) {
	ans := outerTrees([]Point{{1, 1}, {2, 1}, {3, 1}})
	sort.Sort(Points(ans))
	exp := []Point{{1, 1}, {2, 1}, {3, 1}}
	sort.Sort(Points(exp))
	assert.Equal(t, exp, ans)
}

func Test1(t *testing.T) {
	trees := []Point{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}
	exp := []Point{{2, 0}, {4, 2}, {3, 3}, {2, 4}, {1, 1}}
	ans := outerTrees(trees)
	sort.Sort(Points(ans))
	sort.Sort(Points(exp))
	assert.Equal(t, exp, ans)
}

func Test2(t *testing.T) {
	assert.Equal(t, []Point{{1, 1}}, outerTrees([]Point{{1, 1}}))
}
