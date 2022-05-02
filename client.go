package go_tradovate_client

import "net/http"

type (
	baseClient struct {
		*http.Client
	}

	V1Client struct {
		Websocket *WebsocketClient
		REST      *V1RESTClient
	}
)

func NewV1Client() *V1Client {
	client := &http.Client{}

	return &V1Client{
		Websocket: NewWebsocketClient(),
		REST:      NewV1RESTClient(client),
	}
}
