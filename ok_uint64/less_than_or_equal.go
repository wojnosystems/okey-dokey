package ok_uint64

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLessThanOrEqualFormat(definition *LessThanOrEqual, value *uint64) string {
	return fmt.Sprintf("must be less than or equal to %d", definition.Value)
}

type LessThanOrEqual struct {
	Format func(definition *LessThanOrEqual, value *uint64) string
	Value  uint64
}

func (m *LessThanOrEqual) Validate(value *uint64, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLessThanOrEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value > m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
