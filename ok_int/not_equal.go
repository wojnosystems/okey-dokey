package ok_int

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultNotEqualFormat(definition *NotEqual, value int) string {
	return fmt.Sprintf("must not be %d", definition.Value)
}

type NotEqual struct {
	Format func(definition *NotEqual, value int) string
	Value  int
}

func (m *NotEqual) Validate(value optional.Int, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultNotEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual int) {
		if actual == m.Value {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
