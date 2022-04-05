package request_service

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Requester interface {
	SetHeader(header, value string)
}

type GETRequest struct {
	req *http.Request
}
type POSTRequest struct {
	req *http.Request
}

func CreateGETRequest(url string) *GETRequest {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return &GETRequest{req: req}
}

func CreatePOSTRequest(url string, body *bytes.Buffer) *GETRequest {
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return &GETRequest{req: req}
}

func (gr *GETRequest) SetHeader(header, value string) {
	gr.req.Header.Set(header, value)
}

func (pr *POSTRequest) SetHeader(header, value string) {
	pr.req.Header.Set(header, value)
}

func (gr *GETRequest) Send() ([]byte, int, error) {
	client := &http.Client{}

	resp, err := client.Do(gr.req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	return respbody, resp.StatusCode, nil
}

func (pr *POSTRequest) Send() ([]byte, int, error) {
	client := &http.Client{}

	resp, err := client.Do(pr.req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	return respbody, resp.StatusCode, nil
}
