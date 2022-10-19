package govern

import "github.com/pkg/errors"

type (
	StringOpts struct {
		Regex    string
		MinLen   int
		MaxLen   int
		Required bool
	}

	StringType struct {
		required bool
		regex    string
		minLen   int
		maxLen   int
	}
)

func (s StringType) Check(x string) error {
	le := len(x)

	if s.required && le < 1 {
		return ErrRequired
	}

	if s.minLen > 0 && le < s.minLen {
		return errors.Wrapf(ErrMinLength, "%d", s.minLen)
	}

	if s.maxLen > 0 && le > s.minLen {
		return errors.Wrapf(ErrMinLength, "%d", s.minLen)
	}

	return nil
}

func String(opts StringOpts) StringType {
	return StringType{
		required: opts.Required,
		minLen:   opts.MinLen,
		maxLen:   opts.MaxLen,
	}
}
