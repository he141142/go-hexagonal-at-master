package validator

import (
	"errors"
	validator "hex-base/internal/core/adapters/validator"


)


var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

type ValidatorType int

const (
	PLAYGROUND_INSTANCE ValidatorType = iota
)


func NewValidatorFactory(instance ValidatorType) (validator.ValidatorAdapter, error) {
	switch instance {
	case PLAYGROUND_INSTANCE:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}