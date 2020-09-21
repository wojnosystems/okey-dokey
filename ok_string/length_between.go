package ok_string

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthBetweenFormat(definition *LengthBetween, value *string) string {
	return fmt.Sprintf("must have at least %d and at most %d characters, but had %d", definition.AtLeast, definition.AtMost, len(*value))
}

type LengthBetween struct {
	Format  func(definition *LengthBetween, value *string) string
	AtLeast int
	AtMost  int
}

func (m *LengthBetween) Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if len(*value) < m.AtLeast || m.AtMost < len(*value) {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
