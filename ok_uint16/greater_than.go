package ok_uint16

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultGreaterThanFormat(definition *GreaterThan, value *uint16) string {
	return fmt.Sprintf("must be greater than %d", definition.Value)
}

type GreaterThan struct {
	Format func(definition *GreaterThan, value *uint16) string
	Value  uint16
}

func (m *GreaterThan) Validate(value *uint16, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultGreaterThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value <= m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
