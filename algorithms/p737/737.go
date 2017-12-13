package p737

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

	m := make(map[string]string)
	for _, pair := range pairs {
		p1 := find(m, pair[0])
		p2 := find(m, pair[1])

		if p1 != p2 {
			m[p1] = p2
		}
	}

	for i := 0; i < len(words1); i++ {
		if words1[i] != words2[i] && find(m, words1[i]) != find(m, words2[i]) {
			return false
		}
	}
	return true
}

func find(m map[string]string, k string) string {
	_, ok := m[k]
	if !ok {
		m[k] = k
	}
	for m[k] != k {
		m[k] = m[m[k]]
		k = m[k]
	}
	return m[k]
}
