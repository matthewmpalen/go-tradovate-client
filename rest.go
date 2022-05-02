package go_tradovate_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Parameters interface{}
	Data       interface{}
	Response   interface{}

	V1RESTClient struct {
		baseClient
	}

	GetAccessTokenData struct {
		Data
		Name       string `json:"name"`
		Password   string `json:"password"`
		AppID      string `json:"appId"`
		AppVersion string `json:"appVersion"`
		CID        string `json:"cid"`
		DeviceId   string `json:"deviceId"`
		SEC        string `json:"sec"`
	}

	GetAccessTokenResponse struct {
		Response
	}
)

const (
	restDomain = "tradovateapi.com"
)

func NewV1RESTClient(client *http.Client) *V1RESTClient {
	return &V1RESTClient{
		baseClient: baseClient{
			Client: client,
		},
	}
}

func (c V1RESTClient) apiURL() string {
	return fmt.Sprintf("https://%s/v1", restDomain)
}

func (c V1RESTClient) addHeaders(req *http.Request) {
	val := "application/json"
	req.Header.Add("Content-Type", val)
	req.Header.Add("Accept", val)
}

func (c V1RESTClient) decode(req *http.Request, response Response) error {
	c.addHeaders(req)

	resp, doErr := c.baseClient.Do(req)
	if doErr != nil {
		return doErr
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var errorMessage interface{}
		json.NewDecoder(resp.Body).Decode(errorMessage)
		return fmt.Errorf("Invalid status code: %d; resp=%s", resp.StatusCode, errorMessage)
	}

	decodeErr := json.NewDecoder(resp.Body).Decode(response)
	if decodeErr != nil {
		return fmt.Errorf("decoding failed: %s", decodeErr)
	}

	return nil
}

func (c V1RESTClient) request(method, url string, params Parameters, data Data, resp Response) error {
	buf := new(bytes.Buffer)
	encErr := json.NewEncoder(buf).Encode(data)
	if encErr != nil {
		return encErr
	}

	req, reqErr := http.NewRequest(method, url, buf)
	if reqErr != nil {
		return reqErr
	}

	decodeErr := c.decode(req, resp)
	if decodeErr != nil {
		return decodeErr
	}

	return nil
}

func (c V1RESTClient) GetAccessToken(data *GetAccessTokenData) (*GetAccessTokenResponse, error) {
	url := fmt.Sprintf("%s/auth/accesstokenrequest", c.apiURL())

	var resp *GetAccessTokenResponse
	reqErr := c.request(http.MethodPost, url, nil, data, resp)
	if reqErr != nil {
		return nil, reqErr
	}

	return resp, nil
}
