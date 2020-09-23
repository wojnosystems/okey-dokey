package ok_slice

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultItemCountAtMostFormat(definition *ItemCountAtMost, value Counter) string {
	return fmt.Sprintf("cannot have more than %d items", definition.AtMost)
}

type ItemCountAtMost struct {
	Format func(definition *ItemCountAtMost, value Counter) string
	AtMost int
}

func (m *ItemCountAtMost) Validate(value Counter, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountAtMostFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = value.Len()
	}
	if length > m.AtMost {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
