package go_tradovate_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	V1RESTClient struct {
		baseClient
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

func (c V1RESTClient) unmarshal(req *http.Request, response Response) error {
	c.addHeaders(req)

	resp, doErr := c.baseClient.Do(req)
	if doErr != nil {
		return doErr
	}

	defer resp.Body.Close()
	bytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}

	if resp.StatusCode >= http.StatusBadRequest {

		return fmt.Errorf("Invalid status code: %d; resp=%s", resp.StatusCode, string(bytes))
	}

	unmarshalErr := json.Unmarshal(bytes, response)
	if unmarshalErr != nil {
		return fmt.Errorf(
			"unmarshalling failed: %s, status=%d resp=%s",
			unmarshalErr,
			resp.StatusCode,
			string(bytes),
		)
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

	umarshalErr := c.unmarshal(req, resp)
	if umarshalErr != nil {
		return umarshalErr
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

func (c V1RESTClient) PlaceOrder(data *PlaceOrderData) (*PlaceOrderResponse, error) {
	url := fmt.Sprintf("%s/order/placeorder", c.apiURL())

	var resp *PlaceOrderResponse
	reqErr := c.request(http.MethodPost, url, nil, data, resp)
	if reqErr != nil {
		return nil, reqErr
	}

	return resp, nil
}

func (c V1RESTClient) CancelOrder(data *PlaceOrderData) (*CancelOrderResponse, error) {
	url := fmt.Sprintf("%s/order/cancelorder", c.apiURL())

	var resp *CancelOrderResponse
	reqErr := c.request(http.MethodPost, url, nil, data, resp)
	if reqErr != nil {
		return nil, reqErr
	}

	return resp, nil
}

func (c V1RESTClient) PlaceOCO(data *PlaceOCOData) (*PlaceOCOResponse, error) {
	url := fmt.Sprintf("%s/order/placeoco", c.apiURL())

	var resp *PlaceOCOResponse
	reqErr := c.request(http.MethodPost, url, nil, data, resp)
	if reqErr != nil {
		return nil, reqErr
	}

	return resp, nil
}
