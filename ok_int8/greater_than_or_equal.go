package ok_int8

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultGreaterThanOrEqualFormat(definition *GreaterThanOrEqual, value optional.Int8) string {
	return fmt.Sprintf("must be greater than or equal to %d", definition.Value)
}

type GreaterThanOrEqual struct {
	Format func(definition *GreaterThanOrEqual, value optional.Int8) string
	Value  int8
}

func (m *GreaterThanOrEqual) Validate(value optional.Int8, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultGreaterThanOrEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() < m.Value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
