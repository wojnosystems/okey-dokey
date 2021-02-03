package ok_slice_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"strings"
)

func defaultItemsUniqueFormat(definition *ItemsUnique, value []optional.String, duplicateIndexes []int) string {
	sb := strings.Builder{}
	sb.WriteString("must have only unique elements, but had duplicates of ")
	for i, d := range duplicateIndexes {
		value[d].IfSetElse(func(actual string) {
			sb.WriteString(fmt.Sprintf("%v", actual))
		}, func() {
			sb.WriteString("<unset>")
		})
		if i < len(duplicateIndexes)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

type ItemsUnique struct {
	Format func(definition *ItemsUnique, value []optional.String, duplicateIndexes []int) string
}

func (m *ItemsUnique) Validate(value []optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemsUniqueFormat
	if m.Format != nil {
		formatter = m.Format
	}
	unsetCount := 0
	duplicates := make([]int, 0, 10)
	items := make(map[string]bool)
	for i, v := range value {
		v.IfSetElse(func(actual string) {
			if _, ok := items[actual]; ok {
				duplicates = append(duplicates, i)
			} else {
				items[actual] = true
			}
		}, func() {
			unsetCount++
			if unsetCount > 1 {
				duplicates = append(duplicates, i)
			}
		})
	}
	if len(duplicates) != 0 {
		violationReceiver.Emit(formatter(m, value, duplicates))
	}
	return ok_action.Continue
}
