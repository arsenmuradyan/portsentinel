package argocd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
	"jediproj.io/tview-tests/internal"
	"jediproj.io/tview-tests/internal/providers/cd"
)

type ArgoCDProvider struct {
	Configuration internal.ArgoCDConfiguration
}

func (a ArgoCDProvider) GetApplications() []cd.Application {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	body := []byte(fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
	}`, a.Configuration.Username, a.Configuration.Password))
	res, err := http.Post(fmt.Sprintf("%s/api/v1/session", a.Configuration.Url), "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	authResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	token := gjson.Get(string(authResponseBody), "token")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/applications", a.Configuration.Url), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(resp.Body)
	json := string(data)
	applications := gjson.Get(json, "items")
	result := []cd.Application{}
	for _, application := range applications.Array() {
		name := gjson.Get(application.String(), "metadata.name")
		result = append(result, cd.Application(name.Str))
	}
	return result
}
