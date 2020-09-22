package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthAtLeastFormat(definition *LengthAtLeast, value optional.String) string {
	return fmt.Sprintf("cannot have fewer than %d characters", definition.Length)
}

type LengthAtLeast struct {
	Format func(definition *LengthAtLeast, value optional.String) string
	Length int
}

func (m *LengthAtLeast) Validate(value optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if len(value.Value()) < m.Length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
