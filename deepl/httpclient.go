package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RequestBody any

type HTTPClient struct {
	Client  *http.Client
	Headers []Header
	Host    string
}

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Method   string
	Endpoint string
	Body     RequestBody
}

func (c *HTTPClient) NewClient() {
	c.Client = &http.Client{}
	host := os.Getenv("DEEPL_HOST")
	if host[len(host)-1:] == "/" {
		host = host[len(host)-1:]
	}
	c.Host = host

	contentHeader := Header{
		Key:   "Content-Type",
		Value: "application/json",
	}
	authHeader := Header{
		Key:   "Authorization",
		Value: fmt.Sprintf("DeepL-Auth-Key %s", os.Getenv("DEEPL_TOKEN")),
	}
	c.Headers = append(c.Headers, authHeader, contentHeader)
}

func (c *HTTPClient) SendRequest(request Request, target interface{}) error {
	//Build request
	reqBody, err := request.getBody()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", c.Host, request.Endpoint)
	req, err := http.NewRequest(request.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	//Send request using http client
	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("HTTP request error %v", err))
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(resBody) > 0 {
		if err = json.Unmarshal(resBody, target); err != nil {
			return err
		}
	}

	return nil
}

func (r Request) getBody() ([]byte, error) {
	reqBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	return reqBody, nil
}
