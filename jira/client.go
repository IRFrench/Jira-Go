package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"jira-go/cfg"
	"net/http"
	"net/url"
	"time"
)

const (
	searchPath    = "rest/api/3/search"
	authorization = "Authorization"
)

type JiraClient struct {
	client     *http.Client
	authHeader string
	domain     string
}

func (j *JiraClient) RunSearchQuery(jql string) (JiraReponse, error) {
	request, err := j.buildSearchRequest(jql)
	if err != nil {
		return JiraReponse{}, err
	}

	response, err := j.runRequest(request)
	if err != nil {
		return JiraReponse{}, err
	}

	return j.parseResponse(response)
}

func (j *JiraClient) buildSearchRequest(jql string) (*http.Request, error) {
	jqlQuery := url.QueryEscape(jql)
	fullUrl := fmt.Sprintf("%v/%v?jql=%v", j.domain, searchPath, jqlQuery)

	newRequest, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return nil, err
	}

	newRequest.Header.Add(authorization, j.authHeader)
	return newRequest, nil
}

func (j *JiraClient) runRequest(request *http.Request) (*http.Response, error) {
	return j.client.Do(request)
}

func (j *JiraClient) parseResponse(response *http.Response) (JiraReponse, error) {
	parsedResponse := JiraReponse{}

	if err := json.NewDecoder(response.Body).Decode(&parsedResponse); err != nil {
		return JiraReponse{}, err
	}
	return parsedResponse, nil
}

func NewJiraClient(config cfg.Config) JiraClient {
	newClient := http.Client{
		Timeout: 10 * time.Second,
	}
	return JiraClient{
		client:     &newClient,
		authHeader: basicAuth(config.Auth.Email, config.Auth.Token),
		domain:     config.URL,
	}
}

func basicAuth(username string, password string) string {
	usernamePassword := fmt.Sprintf("%v:%v", username, password)
	encodedUsernamePassword := base64.StdEncoding.EncodeToString([]byte(usernamePassword))
	return fmt.Sprintf("Basic %v", encodedUsernamePassword)
}
