package ok_bool

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultFalseFormat(definition *False, value *bool) string {
	return "must be false"
}

type False struct {
	Format func(definition *False, value *bool) string
}

func (m *False) Validate(value *bool, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultFalseFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if *value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
