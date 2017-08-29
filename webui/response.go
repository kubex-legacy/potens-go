package webui

import (
	"bytes"
	"github.com/kubex/proto-go/application"
	"net/http"
	"html/template"
	"encoding/json"
)

//CreateResponse creates a new initialised response
func CreateResponse() *application.HTTPResponse {
	response := &application.HTTPResponse{}
	response.Headers = make(map[string]*application.HTTPResponse_HTTPHeaderParameter)
	return response
}

//CreateJsonResponse Creates a new response
func CreateJsonResponse(content interface{}) *application.HTTPResponse {
	response := CreateResponse()
	response.ContentType = "application/json"
	jsonContent, _ := json.Marshal(content)
	response.Body = string(jsonContent)
	return response
}

//CreateNotModifiedResponse Creates a new 304 response
func CreateNotModifiedResponse() *application.HTTPResponse {
	response := CreateResponse()
	response.StatusCode = http.StatusNotModified
	return response
}

//CreateFromTemplate Create a response from templates using golang html/template
func CreateFromTemplate(data interface{}, filenames ...string) *application.HTTPResponse {
	response := CreateResponse()
	buf := new(bytes.Buffer)
	template.Must(template.ParseFiles(filenames...)).Execute(buf, data)
	response.Body = buf.String()
	return response
}
