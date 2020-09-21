package ok_uint64

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLessThanFormat(definition *LessThan, value *uint64) string {
	return fmt.Sprintf("must be less than %d", definition.Value)
}

type LessThan struct {
	Format func(definition *LessThan, value *uint64) string
	Value  uint64
}

func (m *LessThan) Validate(value *uint64, violationReceiver bad.MessageReceiver) ok_action.Enum {
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
