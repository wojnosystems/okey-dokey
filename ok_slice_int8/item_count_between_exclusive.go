package ok_slice_int8

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"github.com/wojnosystems/okey-dokey/ok_range"
)

func defaultItemCountBetweenExclusiveFormat(definition *ItemCountBetweenExclusive, value []optional.Int8) string {
	return fmt.Sprintf("must have more than %d but fewer than %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetweenExclusive struct {
	Format  func(definition *ItemCountBetweenExclusive, value []optional.Int8) string
	Between ok_range.IntExclusive
}

func (m *ItemCountBetweenExclusive) Validate(value []optional.Int8, violationReceiver bad.Emitter) ok_action.Enum {
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
