package errors

import "fmt"

// Errors model
type Errors struct {
	Name    string      `json:"name,omitempty"`
	Message string      `json:"message,omitempty"`
	Panic   bool        `json:"panic,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e Errors) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

func New(name, message string) Errors {
	return Errors{
		Name:    name,
		Message: message,
	}
}

func NewError(name string, e error) Errors {
	return Errors{
		Name:    name,
		Message: e.Error(),
	}
}
