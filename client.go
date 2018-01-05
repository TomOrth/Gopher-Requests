//Package gopherReq provides a wrapper on the HTTP client
package gopherReq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Type Client represents the HTTP client that contains a base url
type Client struct {
	Base string
}

var (
	httpclient *http.Client
	req        *http.Request
)

//req intializes the default http client and creates a new request
func (client Client) req(method, path string) (*http.Request, error) {
	httpclient = &http.Client{}
	url := client.Base + path
	return http.NewRequest(method, url, nil)
}

//Get creates a GET request at the given path on the base url of the client
func (client Client) Get(path string) Client {
	req, _ = client.req("GET", path)
	return client
}

//Put creates a PUT request at the given path on the base url of the client
func (client Client) Put(path string) Client {
	req, _ = client.req("PUT", path)
	return client
}

//Post creates a POST request at the given path on the base url of the client
func (client Client) Post(path string) Client {
	req, _ = client.req("POST", path)
	return client
}

//Patch creates a PATCH request at the given path on the base url of the client
func (client Client) Patch(path string) Client {
	req, _ = client.req("PATCH", path)
	return client
}

//Delete creates a DELETE request at the given path on the base url of the client
func (client Client) Delete(path string) Client {
	req, _ = client.req("DELETE", path)
	return client
}

//Header addes header to the current request
func (client Client) Header(key string, value string) Client {
	req.Header.Add(key, value)
	return client
}

//Exec executes the request and writes the results to v. If an issue occurs, an error is returned
func (client Client) Exec(v interface{}) error {
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("non-200 response code %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, errResp := ioutil.ReadAll(resp.Body)

	if errResp != nil {
		return errResp
	}
	errF := json.Unmarshal(body, v)
	return errF
}
