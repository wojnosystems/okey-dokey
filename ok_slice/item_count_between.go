package ok_slice

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"okey-dokey/ok_range"
)

func defaultItemCountBetweenFormat(definition *ItemCountBetween, value Counter) string {
	return fmt.Sprintf("must have between %d and %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetween struct {
	Format  func(definition *ItemCountBetween, value Counter) string
	Between ok_range.Int
}

func (m *ItemCountBetween) Validate(value Counter, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = value.Len()
	}
	if length < m.Between.Start() || m.Between.End() < length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
