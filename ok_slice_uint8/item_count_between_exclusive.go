package ok_slice_uint8

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"okey-dokey/ok_range"
)

func defaultItemCountBetweenExclusiveFormat(definition *ItemCountBetweenExclusive, value []optional.Uint8) string {
	return fmt.Sprintf("must have more than %d but fewer than %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetweenExclusive struct {
	Format  func(definition *ItemCountBetweenExclusive, value []optional.Uint8) string
	Between ok_range.IntExclusive
}

func (m *ItemCountBetweenExclusive) Validate(value []optional.Uint8, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemCountBetweenExclusiveFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length <= m.Between.Start() || m.Between.End() <= length {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
