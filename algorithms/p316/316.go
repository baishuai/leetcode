package p316

/**
Given a string which contains only lowercase letters,
remove duplicate letters so that every letter appear once and only once.
You must make sure your result is the smallest in lexicographical order
among all possible results.

Example:
Given "bcabc"
Return "abc"

Given "cbacdcbc"
Return "acdb"
*/

func removeDuplicateLetters(s string) string {

	bs := []byte(s)
	cnt := make([]int, 26)
	visited := make([]bool, 26)
	for _, v := range bs {
		cnt[v-'a'] ++
	}
	res := make([]byte, 0)
	res = append(res, '-')
	for _, v := range bs {
		cnt[v-'a']--
		if visited[v-'a'] {
			continue
		}

		back := res[len(res)-1]
		for back > v && cnt[back-'a'] > 0 {
			visited[back-'a'] = false
			res = res[:len(res)-1]
			back = res[len(res)-1]
		}
		res = append(res, v)
		visited[v-'a'] = true
	}
	return string(res[1:])
}
