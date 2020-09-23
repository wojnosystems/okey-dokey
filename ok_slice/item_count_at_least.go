package ok_slice

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultItemCountAtLeastFormat(definition *ItemCountAtLeast, value Counter) string {
	return fmt.Sprintf("cannot have fewer than %d items", definition.AtLeast)
}

type ItemCountAtLeast struct {
	Format  func(definition *ItemCountAtLeast, value Counter) string
	AtLeast int
}

func (m *ItemCountAtLeast) Validate(value Counter, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = value.Len()
	}
	if length < m.AtLeast {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
