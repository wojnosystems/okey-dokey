package ok_int8

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLessThanFormat(definition *LessThan, value *int8) string {
	return fmt.Sprintf("must be less than %d", definition.Value)
}

type LessThan struct {
	Format func(definition *LessThan, value *int8) string
	Value  int8
}

func (m *LessThan) Validate(value *int8, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLessThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value >= m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
