package forms

import "net/http"

type FormResponse struct {
	WasSuccessful bool
	Errors        map[string][]string
	ToastMessage  string
	RedirectPath  string
	StatusCode    int
}

func NewResponse() *FormResponse {
	frmResp := &FormResponse{}
	frmResp.Errors = make(map[string][]string)
	frmResp.StatusCode = http.StatusOK
	frmResp.WasSuccessful = true
	return frmResp
}

func (r *FormResponse) AddError(field string, err string) {
	if _, ok := r.Errors[field]; !ok {
		r.Errors[field] = []string{}
	}

	r.Errors[field] = append(r.Errors[field], err)
	r.WasSuccessful = false
}
