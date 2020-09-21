package ok_uint

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultEqualFormat(definition *Equal, value *uint) string {
	return fmt.Sprintf("must be exactly %d", definition.Value)
}

type Equal struct {
	Format func(definition *Equal, value *uint) string
	Value  uint
}

func (m *Equal) Validate(value *uint, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value != m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
