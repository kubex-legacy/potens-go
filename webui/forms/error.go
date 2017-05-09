package forms

type FormError struct {
	Errors map[string][]string
}

//NewError New form error handler
func NewError() *FormError {
	frmErr := &FormError{}
	frmErr.Errors = make(map[string][]string)
	return frmErr
}

func (e *FormError) AddError(field string, err string) {
	if _, ok := e.Errors[field]; !ok {
		e.Errors[field] = []string{}
	}

	e.Errors[field] = append(e.Errors[field], err)
}
