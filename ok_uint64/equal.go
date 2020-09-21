package ok_uint64

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultEqualFormat(definition *Equal, value *uint64) string {
	return fmt.Sprintf("must be exactly %d", definition.Value)
}

type Equal struct {
	Format func(definition *Equal, value *uint64) string
	Value  uint64
}

func (m *Equal) Validate(value *uint64, violationReceiver bad.MessageReceiver) ok_action.Enum {
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
