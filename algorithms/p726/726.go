package p726

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
	"strings"
)

type Atom struct {
	atom string
	cnt  int
}

type Atoms []Atom

// Len is the number of elements in the collection.
func (a Atoms) Len() int {
	return len(a)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (a Atoms) Less(i, j int) bool {
	return a[i].atom < a[j].atom
}

// Swap swaps the elements with indexes i and j.
func (a Atoms) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func parseFormula(reader *bufio.Reader) map[string]int {
	counts := make(map[string]int)

	b, err := reader.ReadByte()

	for err == nil && b != ')' {
		reader.UnreadByte()
		cnts := parseUnit(reader)

		for k, v := range cnts {
			counts[k] += v
		}
		b, err = reader.ReadByte()
	}

	if err == nil && b != ')' {
		reader.UnreadByte()
	}
	return counts
}

func parseUnit(reader *bufio.Reader) map[string]int {
	counts := make(map[string]int)
	b, err := reader.ReadByte()
	if err == nil && b == '(' {
		cnts := parseFormula(reader)

		digits := parseInt(reader)
		for k, v := range cnts {
			counts[k] += v * digits
		}
	} else if err == nil {
		buf := new(bytes.Buffer)
		buf.WriteByte(b)

		b, err = reader.ReadByte()
		for err == nil && b >= 'a' && b <= 'z' {
			buf.WriteByte(b)
			b, err = reader.ReadByte()
		}
		if err == nil {
			reader.UnreadByte()
		}
		atom := buf.String()
		digits := parseInt(reader)

		counts[atom] += digits
	}
	return counts
}

func parseInt(reader *bufio.Reader) int {
	b, err := reader.ReadByte()

	val := 0
	for err == nil && b <= '9' && b >= '0' {
		val = val*10 + int(b-'0')
		b, err = reader.ReadByte()
	}
	if val == 0 {
		val = 1
	}

	if err == nil {
		reader.UnreadByte()
	}
	return val
}

func countOfAtoms(formula string) string {

	reader := bufio.NewReader(strings.NewReader(formula))

	counts := parseFormula(reader)
	atoms := make([]Atom, 0)
	for k, v := range counts {
		atoms = append(atoms, Atom{k, v})
	}

	sort.Sort(Atoms(atoms))

	var buffer bytes.Buffer

	for _, atom := range atoms {
		buffer.WriteString(atom.atom)
		if atom.cnt > 1 {
			buffer.WriteString(strconv.Itoa(atom.cnt))
		}
	}

	return buffer.String()
}
