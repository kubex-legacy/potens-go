package forms

import "net/http"

type FormResponse struct {
	Errors       map[string][]string
	ToastMessage string
	RedirectPath string
	StatusCode   int
}

func NewResponse() *FormResponse {
	frmResp := &FormResponse{}
	frmResp.Errors = make(map[string][]string)
	frmResp.StatusCode = http.StatusOK
	return frmResp
}
