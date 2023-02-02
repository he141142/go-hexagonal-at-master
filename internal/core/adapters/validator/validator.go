package validator

type ValidatorAdapter interface {
	Validate(interface{}) error
	Messages() []string
}
