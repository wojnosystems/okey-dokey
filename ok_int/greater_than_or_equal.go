package ok_int

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultGreaterThanOrEqualFormat(definition *GreaterThanOrEqual, value *int) string {
	return fmt.Sprintf("must be greater than or equal to %d", definition.Value)
}

type GreaterThanOrEqual struct {
	Format func(definition *GreaterThanOrEqual, value *int) string
	Value  int
}

func (m *GreaterThanOrEqual) Validate(value *int, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultGreaterThanOrEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value < m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
