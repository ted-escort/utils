package utils

type ErrorUtil struct {
	Message string
}

func NewErrorUtil(message string) *ErrorUtil {
	return &ErrorUtil{Message: message}
}

func (this *ErrorUtil) Error() string {
	return this.Message
}
