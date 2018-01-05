//Package gopherReq provides a wrapper on the HTTP client
package gopherReq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Type Client represents the HTTP client
type Client struct {
	Base string
}

var (
	httpclient *http.Client
	req        *http.Request
)

func (client Client) req(method, path string) (*http.Request, error) {
	httpclient = &http.Client{}
	url := client.Base + path
	return http.NewRequest(method, url, nil)
}

func (client Client) Get(path string) Client {
	req, _ = client.req("GET", path)
	return client
}

func (client Client) Put(path string) Client {
	req, _ = client.req("PUT", path)
	return client
}

func (client Client) Post(path string) Client {
	req, _ = client.req("POST", path)
	return client
}

func (client Client) Patch(path string) Client {
	req, _ = client.req("PATCH", path)
	return client
}

func (client Client) Delete(path string) Client {
	req, _ = client.req("DELETE", path)
	return client
}

func (client Client) Header(key string, value string) Client {
	req.Header.Add(key, value)
	return client
}

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
