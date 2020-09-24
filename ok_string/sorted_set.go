package ok_string

import "sort"

type SortedSet sort.StringSlice

type SortedSetBuilder struct {
	values map[string]bool
}

func NewSortedSetBuilder(values ...string) (b SortedSetBuilder) {
	for _, value := range values {
		b.Add(value)
	}
	return
}

func (b *SortedSetBuilder) Add(value string) *SortedSetBuilder {
	if b.values == nil {
		b.values = make(map[string]bool)
	}
	b.values[value] = true
	return b
}

func (b SortedSetBuilder) Build() (out SortedSet) {
	out = make([]string, len(b.values))
	index := 0
	for key := range b.values {
		out[index] = key
		index++
	}
	sort.Strings(out)
	return
}
