package egerror

type Error struct {
	E       error
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.E.Error()
}
