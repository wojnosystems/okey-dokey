package ok_uint8

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultNotEqualFormat(definition *NotEqual, value *uint8) string {
	return fmt.Sprintf("must not be %d", definition.Value)
}

type NotEqual struct {
	Format func(definition *NotEqual, value *uint8) string
	Value  uint8
}

func (m *NotEqual) Validate(value *uint8, violationReceiver bad.MessageReceiver) ok_action.Enum {
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
