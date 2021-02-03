package ok_uint8

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultGreaterThanOrEqualFormat(definition *GreaterThanOrEqual, value uint8) string {
	return fmt.Sprintf("must be greater than or equal to %d", definition.Value)
}

type GreaterThanOrEqual struct {
	Format func(definition *GreaterThanOrEqual, value uint8) string
	Value  uint8
}

func (m *GreaterThanOrEqual) Validate(value optional.Uint8, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultGreaterThanOrEqualFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual uint8) {
		if actual < m.Value {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
