package health

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpStatusChecker struct {
	BaseUrl string
	Name    string
}

func (h HttpStatusChecker) Traverse(traversalPath []string, action string) (string, error) {
	dependencies := ""
	if len(traversalPath) > 0 {
		dependencies = fmt.Sprintf("&dependencies=%s", strings.Join(traversalPath, ","))
	}
	// Build our HTTP request
	url := fmt.Sprintf("%s/status/traverse?action=%s%s", h.BaseUrl, action, dependencies)
	req, err := http.NewRequest("GET", url, nil)
	// req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s \n", err.Error())
		return "", err
	}
	// Execute HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %s \n", err.Error())
		return "", err
	}
	// Defer the closing of the body
	defer resp.Body.Close()
	// Read our response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err.Error())
		return "", err
	}
	// Return our response
	return string(responseBody), nil

}
