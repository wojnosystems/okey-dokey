package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	sorted_set "github.com/wojnosystems/go-sorted-set"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"sort"
	"strings"
)

const (
	oneOfMaxItemsToList = 10
)

func defaultFormatOneOf(definition *OneOf, value string) string {
	itemsToList := oneOfMaxItemsToList
	if len(definition.Only) < itemsToList {
		itemsToList = len(definition.Only)
	}
	list := strings.Join(definition.Only[0:itemsToList], ", ")
	if itemsToList < len(definition.Only) {
		list += ", ..."
	}
	return fmt.Sprintf("must be one of the following: %s", list)
}

type OneOf struct {
	Format func(definition *OneOf, value string) string
	Only   sorted_set.String
}

func (m *OneOf) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultFormatOneOf
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		searchIndex := sort.SearchStrings(m.Only, actual)
		if searchIndex == len(m.Only) || m.Only[searchIndex] != actual {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
