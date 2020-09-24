package ok_slice_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"okey-dokey/ok_range"
)

func defaultItemCountBetweenFormat(definition *ItemCountBetween, value []optional.String) string {
	return fmt.Sprintf("must have between %d and %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetween struct {
	Format  func(definition *ItemCountBetween, value []optional.String) string
	Between ok_range.Int
}

func (m *ItemCountBetween) Validate(value []optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length < m.Between.Start() || m.Between.End() < length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}