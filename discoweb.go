package potens

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type discoverWebService struct {
	Host      string `json:"host"`
	Address   []string `json:"addresses"`
	Available bool`json:"available"`
}

func (app *Application) getCoreService(serviceName string) (*discoverWebService, error) {
	discoWebUrl := os.Getenv("KDISCO_WEB_URL")
	if discoWebUrl == "" {
		discoWebUrl = "http://disco." + app.KubexDomain()
	}

	response, err := http.Get(discoWebUrl + "/" + serviceName)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	service := &discoverWebService{}
	err = json.Unmarshal(buf.Bytes(), service)
	if err != nil {
		return nil, err
	}

	return service, nil
}
