package ok_string

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"net/url"
)

func defaultIsURL(definition *IsURL, value optional.String) string {
	return "is not a URL"
}

type IsURL struct {
	Format func(definition *IsURL, value optional.String) string
}

func (m *IsURL) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultIsURL
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		_, err := url.Parse(actual)
		if err != nil {
			violationReceiver.Emit(formatter(m, value))
		}
	})
	return ok_action.Continue
}
