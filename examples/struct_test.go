package examples

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_int"
	"okey-dokey/ok_range"
	"okey-dokey/ok_slice_string"
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
	IceCreamFlavors ok_slice_string.On
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
	IceCreamFlavors: ok_slice_string.On{
		Id: "iceCreamFlavors",
		Ensure: []ok_slice_string.Definer{
			&ok_slice_string.ItemCountBetween{
				Between: ok_range.IntBetween(2, 10),
			},
		},
		EnsureItems: []ok_string.Definer{
			&ok_string.OneOf{
				Only: ok_string.NewSortedSetBuilder("raspberry", "vanilla", "chocolate", "pistachio").
					Build(),
			},
		},
	},
}

// This method should be provided by the library as a reflection implementation
// It just goes through all of the items and evaluates them, then finally, evaluates the entire struct
func (v modelValidationDefs) Validate(on *model, receiver bad.MemberReceiver) {
	ok_string.Validate(on.Name, &v.Name, receiver)
	ok_int.Validate(on.Age, &v.Age, receiver)
	ok_slice_string.Validate(on.IceCreamFlavors, &v.IceCreamFlavors, receiver)
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
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{},
		},
		"string missing": {
			m: model{
				Age: optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{
				"name": {"is required"},
			},
		},
		"age missing": {
			m: model{
				Name: optional.StringFrom("chris"),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{
				"age": {"is required"},
			},
		},
		"age too young": {
			m: model{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(17),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{
				"age": {"must be greater than or equal to 18"},
			},
		},
		"name too long": {
			m: model{
				Name: optional.StringFrom("veryLongName"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("chocolate"),
					optional.StringFrom("vanilla"),
				},
			},
			expected: map[string][]string{
				"name": {"cannot have more than 10 characters"},
			},
		},
		"bad ice cream": {
			m: model{
				Name: optional.StringFrom("chris"),
				Age:  optional.IntFrom(30),
				IceCreamFlavors: []optional.String{
					optional.StringFrom("coffee"),
					optional.StringFrom("chocolate"),
					optional.StringFrom("orange"),
				},
			},
			expected: map[string][]string{
				"iceCreamFlavors[0]": {"must be one of the following: chocolate, pistachio, raspberry, vanilla"},
				"iceCreamFlavors[2]": {"must be one of the following: chocolate, pistachio, raspberry, vanilla"},
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
