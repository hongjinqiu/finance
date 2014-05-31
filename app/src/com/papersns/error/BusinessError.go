package error

import (
)

type BusinessError struct{
	Message string
}

func (e *BusinessError) Error() string {
	return e.Message
}
