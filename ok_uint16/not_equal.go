package ok_uint16

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultNotEqualFormat(definition *NotEqual, value *uint16) string {
	return fmt.Sprintf("must not be %d", definition.Value)
}

type NotEqual struct {
	Format func(definition *NotEqual, value *uint16) string
	Value  uint16
}

func (m *NotEqual) Validate(value *uint16, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultNotEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value == m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
