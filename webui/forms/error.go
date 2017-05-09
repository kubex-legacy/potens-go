package forms

func (r *FormResponse) AddError(field string, err string) {
	if _, ok := r.Errors[field]; !ok {
		r.Errors[field] = []string{}
	}

	r.Errors[field] = append(r.Errors[field], err)
}
