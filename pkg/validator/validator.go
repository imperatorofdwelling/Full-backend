package validator

import (
	"regexp"
)

var (
	EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
)

type Validator struct {
	Errors map[string]string
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, msg string) {
	v.Errors[key] = msg
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
