package ok_slice

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"okey-dokey/ok_range"
)

func defaultItemCountBetweenExclusiveFormat(definition *ItemCountBetweenExclusive, value Counter) string {
	return fmt.Sprintf("must have more than %d but fewer than %d items", definition.Between.Start(), definition.Between.End())
}

type ItemCountBetweenExclusive struct {
	Format  func(definition *ItemCountBetweenExclusive, value Counter) string
	Between ok_range.IntExclusive
}

func (m *ItemCountBetweenExclusive) Validate(value Counter, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountBetweenExclusiveFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = value.Len()
	}
	if length <= m.Between.Start() || m.Between.End() <= length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
