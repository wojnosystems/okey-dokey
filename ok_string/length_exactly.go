package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultFormatLengthExactly(definition *LengthExactly, value string) string {
	return fmt.Sprintf("was not exactly %d characters", definition.Length)
}

type LengthExactly struct {
	Format func(definition *LengthExactly, value string) string
	Length int
}

func (m *LengthExactly) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultFormatLengthExactly
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		if len(actual) != m.Length {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
