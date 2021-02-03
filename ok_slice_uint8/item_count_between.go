package ok_slice_uint8

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"github.com/wojnosystems/okey-dokey/ok_range"
)

func defaultItemCountBetweenFormat(definition *ItemCountBetween, value []optional.Uint8) string {
	return fmt.Sprintf("must have between %d and %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetween struct {
	Format  func(definition *ItemCountBetween, value []optional.Uint8) string
	Between ok_range.Int
}

func (m *ItemCountBetween) Validate(value []optional.Uint8, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemCountBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length < m.Between.Start() || m.Between.End() < length {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
