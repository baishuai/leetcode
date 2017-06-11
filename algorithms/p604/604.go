package p604

/**
Design and implement a data structure for a compressed string iterator. It should support the following operations: next and hasNext.

The given compressed string will be in the form of each letter followed by a positive integer representing the number of this letter existing in the original uncompressed string.

next() - if the original string still has uncompressed characters, return the next letter; Otherwise return a white space.
hasNext() - Judge whether there is any letter needs to be uncompressed.
*/

type StringIterator struct {
	bs  []byte
	cnt []int
}

func Constructor(compressedString string) StringIterator {
	n := len(compressedString)
	bs := make([]byte, 0, n/2)
	cnt := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		bs = append(bs, compressedString[i])
		i++
		num := 0
		for i < n && compressedString[i] >= '0' && compressedString[i] <= '9' {
			num = num*10 + int(compressedString[i]-'0')
			i++
		}
		i--
		cnt = append(cnt, num)
	}
	return StringIterator{bs, cnt}
}

func (this *StringIterator) Next() byte {
	if len(this.bs) > 0 {
		t := this.bs[0]
		this.cnt[0]--
		if this.cnt[0] == 0 {
			this.cnt = this.cnt[1:]
			this.bs = this.bs[1:]
		}
		return t
	}
	return ' '
}

func (this *StringIterator) HasNext() bool {
	return len(this.bs) > 0
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
