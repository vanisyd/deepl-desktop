package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type RequestBody interface{}

type HTTPClient struct {
	Headers []Header
	Client http.Client
}

type Header struct {
	Key string
	Value string
}

type Request struct {
	Method string
	Endpoint string
	Body RequestBody
}

func (c *HTTPClient) NewClient() {
	c.Client = http.Client{}

	host := os.Getenv("DEEPL_HOST")
	if host[len(host)-1:] == "/" {
		host = host[len(host)-1:]
	}
	hostHeader := Header{
		Key:   "Host",
		Value: host,
	}
	authHeader := Header{
		Key:   "Authorization",
		Value: fmt.Sprintf("DeepL-Auth-Key %s", os.Getenv("DEEPL_TOKEN")),
	}
	c.Headers = append(c.Headers, authHeader, hostHeader)
}

func (c *HTTPClient) SendRequest(request Request) ([]byte, error) {
	//Build request
	reqBody, err := request.getBody()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(request.Method, "", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	for _,header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	//Send request using http client
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Status != strconv.Itoa(http.StatusOK) {
		return nil, errors.New(fmt.Sprintf("Request error: %s", res.Status))
	}

	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)

	return resBody, nil
}

func (r Request) getBody() ([]byte, error) {
	reqBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}


	return reqBody, nil
}
