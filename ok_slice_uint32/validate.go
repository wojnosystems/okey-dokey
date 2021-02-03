package ok_slice_uint32

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"strconv"
)

func Validate(values []optional.Uint32, on *On, violations bad.MemberEmitter) {
	for _, definition := range on.Ensure {
		action := definition.Validate(values, violations)
		if action != ok_action.Continue {
			return
		}
	}
	for i, value := range values {
		mem := violations.Into("[" + strconv.FormatInt(int64(i), 10) + "]")
		for _, definition := range on.EnsureItems {
			action := definition.Validate(value, mem)
			if action != ok_action.Continue {
				return
			}
		}
	}
}
