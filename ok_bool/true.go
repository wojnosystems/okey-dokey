package ok_bool

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultTrueFormat(definition *True, value *bool) string {
	return "must be true"
}

type True struct {
	Format func(definition *True, value *bool) string
}

func (m *True) Validate(value *bool, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultTrueFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if !*value {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
