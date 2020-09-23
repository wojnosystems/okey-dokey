package examples

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_int"
	"okey-dokey/ok_slice"
	"okey-dokey/ok_string"
	"testing"
)

type model struct {
	Name            optional.String
	Age             optional.Int
	IceCreamFlavors []optional.String
}

type modelValidationDefs struct {
	Name            ok_string.On
	Age             ok_int.On
	IceCreamFlavors ok_slice.On
	IceCreamFlavor  ok_string.On
}

// Validations defined separately so that they can be swapped out on the model, depending on the situation
// This pre-compiles the validations so that they exist only once. Then the validations, which have no state,
// are run on the model
var modelValidations = modelValidationDefs{
	Name: ok_string.On{
		Id: "name",
		Ensure: []ok_string.Definer{
			&ok_string.IsRequired{},
			&ok_string.LengthAtMost{
				Length: 10,
			},
		},
	},
	Age: ok_int.On{
		Id: "age",
		Ensure: []ok_int.Definer{
			&ok_int.IsRequired{},
			&ok_int.GreaterThanOrEqual{
				Value: 18,
			},
		},
	},
	IceCreamFlavors: ok_slice.On{
		Id: "icecream_flavors",
		Ensure: []ok_slice.Definer{
			&ok_slice.ItemCountAtLeast{
				AtLeast: 2,
			},
		},
	},
}

// This method should be provided by the library as a reflection implementation
// It just goes through all of the items and evaluates them, then finally, evaluates the entire struct
func (v modelValidationDefs) Validate(on *model, receiver bad.MemberReceiver) {
	ok_string.Validate(on.Name, &v.Name, receiver)
	ok_int.Validate(on.Age, &v.Age, receiver)
}

func TestModel(t *testing.T) {
	cases := map[string]struct {
		m        model
		expected map[string][]string
	}{
		"ok": {
			m: model{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(30),
			},
			expected: map[string][]string{},
		},
		"string missing": {
			m: model{
				Age: optional.IntFrom(30),
			},
			expected: map[string][]string{
				"name": {"is required"},
			},
		},
		"age missing": {
			m: model{
				Name: optional.StringFrom("chris"),
			},
			expected: map[string][]string{
				"age": {"is required"},
			},
		},
		"age too young": {
			m: model{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(17),
			},
			expected: map[string][]string{
				"age": {"must be greater than or equal to 18"},
			},
		},
		"name too long": {
			m: model{
				Name: optional.StringFrom("chriswojno1"),
				Age:  optional.IntFrom(30),
			},
			expected: map[string][]string{
				"name": {"cannot have more than 10 characters"},
			},
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := bad.SliceMemberReceiver{}
			modelValidations.Validate(&c.m, &actual)
			assert.Equal(t, bad.SliceMemberReceiver(c.expected), actual)
		})
	}
}
