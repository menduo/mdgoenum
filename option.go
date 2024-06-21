// Created by @menduo @ 2024/6/21
// a placholder
package mdgoenum

// optionType is the type of the options.
type optionType struct {
	strict bool // strict mode
}

// OpsFuncType is the type of the option functions.
type OpsFuncType func(op *optionType)

// newOption creates a new option.
func newOption() *optionType {
	return &optionType{
		strict: false,
	}
}

// newOptionWithOpts creates a new option with the given options.
func newOptionWithOpts(opFuncs ...OpsFuncType) *optionType {
	op := newOption()
	op.strict = false

	for _, opt := range opFuncs {
		opt(op)
	}

	return op
}
