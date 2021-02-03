package ok_slice_uint32

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"github.com/wojnosystems/okey-dokey/ok_range"
)

func defaultItemCountBetweenExclusiveFormat(definition *ItemCountBetweenExclusive, value []optional.Uint32) string {
	return fmt.Sprintf("must have more than %d but fewer than %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetweenExclusive struct {
	Format  func(definition *ItemCountBetweenExclusive, value []optional.Uint32) string
	Between ok_range.IntExclusive
}

func (m *ItemCountBetweenExclusive) Validate(value []optional.Uint32, violationReceiver bad.Emitter) ok_action.Enum {
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
