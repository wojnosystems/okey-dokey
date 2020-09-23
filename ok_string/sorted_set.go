package ok_string

import "sort"

type SortedSet sort.StringSlice

type SortedSetBuilder struct {
	values map[string]bool
}

func (b *SortedSetBuilder) Add(value string) *SortedSetBuilder {
	if b.values == nil {
		b.values = make(map[string]bool)
	}
	b.values[value] = true
	return b
}

func (b SortedSetBuilder) Sort() (out SortedSet) {
	out = make([]string, len(b.values))
	index := 0
	for key, _ := range b.values {
		out[index] = key
		index++
	}
	sort.Strings(out)
	return
}
