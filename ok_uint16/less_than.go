package ok_uint16

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultLessThanFormat(definition *LessThan, value uint16) string {
	return fmt.Sprintf("must be less than %d", definition.Value)
}

type LessThan struct {
	Format func(definition *LessThan, value uint16) string
	Value  uint16
}

func (m *LessThan) Validate(value optional.Uint16, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultLessThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual uint16) {
		if actual >= m.Value {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
