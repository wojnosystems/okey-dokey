package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultFormatLengthExactly(definition *LengthExactly, value optional.String) string {
	return fmt.Sprintf("was not exactly %d characters", definition.Length)
}

type LengthExactly struct {
	Format func(definition *LengthExactly, value optional.String) string
	Length int
}

func (m *LengthExactly) Validate(value optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultFormatLengthExactly
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if len(value.Value()) != m.Length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
