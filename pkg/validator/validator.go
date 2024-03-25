package validator

type NullFieldValidator interface {
	TableExists() bool
	ValidateNullFields() ([]interface{}, error)
}
