package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultFalseFormat(definition *False, value optional.Bool) string {
	return "must be false"
}

type False struct {
	Format func(definition *False, value optional.Bool) string
}

func (m *False) Validate(value optional.Bool, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultFalseFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if value.Value() {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
