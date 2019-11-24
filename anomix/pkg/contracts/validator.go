package contracts

//Validator
type Validator interface {
	Validate() *Error
}
