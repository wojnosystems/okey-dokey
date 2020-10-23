package ok_slice_uint

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"strings"
)

func defaultItemsUniqueFormat(definition *ItemsUnique, value []optional.Uint, duplicateIndexes []int) string {
	sb := strings.Builder{}
	sb.WriteString("must have only unique elements, but had duplicates of ")
	for i, d := range duplicateIndexes {
		sb.WriteString(fmt.Sprintf("%v", value[d].Value()))
		if i < len(duplicateIndexes)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

type ItemsUnique struct {
	Format func(definition *ItemsUnique, value []optional.Uint, duplicateIndexes []int) string
}

func (m *ItemsUnique) Validate(value []optional.Uint, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemsUniqueFormat
	if m.Format != nil {
		formatter = m.Format
	}
	duplicates := make([]int, 0, 10)
	items := make(map[uint]bool)
	for i, v := range value {
		if _, ok := items[v.Value()]; ok {
			duplicates = append(duplicates, i)
		} else {
			items[v.Value()] = true
		}
	}
	if len(duplicates) != 0 {
		violationReceiver.Emit(formatter(m, value, duplicates))
	}
	return ok_action.Continue
}
