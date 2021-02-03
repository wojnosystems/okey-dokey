package ok_uint32

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultGreaterThanFormat(definition *GreaterThan, value uint32) string {
	return fmt.Sprintf("must be greater than %d", definition.Value)
}

type GreaterThan struct {
	Format func(definition *GreaterThan, value uint32) string
	Value  uint32
}

func (m *GreaterThan) Validate(value optional.Uint32, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultGreaterThanFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual uint32) {
		if actual <= m.Value {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
