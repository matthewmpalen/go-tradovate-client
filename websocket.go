package go_tradovate_client

type (
	WebsocketClient struct {
		//dialer ws.Dialer
	}
)

func NewWebsocketClient() *WebsocketClient {
	return &WebsocketClient{
		//dialer: ws.Dialer{},
	}
}
