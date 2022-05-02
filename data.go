package go_tradovate_client

import "time"

type (
	Data interface{}

	GetAccessTokenData struct {
		Data
		Name       string `json:"name"`
		Password   string `json:"password"`
		AppID      string `json:"appId,omitempty"`
		AppVersion string `json:"appVersion,omitempty"`
		CID        string `json:"cid,omitempty"`
		DeviceId   string `json:"deviceId,omitempty"`
		SEC        string `json:"sec,omitempty"`
	}

	PlaceOrderData struct {
		Data
		AccountSpec    string    `json:"accountSpec"`
		AccountID      int       `json:"accountId"`
		ClOrdID        string    `json:"clOrdId"`
		Action         string    `json:"action"`
		Symbol         string    `json:"symbol"`
		OrderQty       int       `json:"orderQty"`
		OrderType      string    `json:"orderType"`
		Price          int       `json:"price"`
		StopPrice      int       `json:"stopPrice,omitempty"`
		MaxShow        int       `json:"maxShow"`
		PegDifference  int       `json:"pegDifference"`
		TimeInForce    string    `json:"timeInForce"`
		ExpireTime     time.Time `json:"expireTime,omitempty"`
		Text           string    `json:"text,omitempty"`
		ActivationTime time.Time `json:"activationTime,omitempty"`
		CustomTag50    string    `json:"customTag50,omitempty"`
		IsAutomated    bool      `json:"IsAutomated"`
	}

	CancelOrderData struct {
		Data
		OrderID        int       `json:"orderId"`
		ClOrdID        string    `json:"clOrdId"`
		ActivationTime time.Time `json:"activationTime,omitempty"`
		CustomTag50    string    `json:"customTag50,omitempty"`
		IsAutomated    bool      `json:"isAutomated"`
	}

	PlaceOCOOther struct {
		Action        string    `json:"action"`
		ClOrdID       string    `json:"clOrdId"`
		OrderType     string    `json:"orderType"`
		Price         int       `json:"price"`
		StopPrice     int       `json:"stopPrice,omitempty"`
		MaxShow       int       `json:"maxShow"`
		PegDifference int       `json:"pegDifference"`
		TimeInForce   string    `json:"timeInForce"`
		ExpireTime    time.Time `json:"expireTime,omitempty"`
		Text          string    `json:"text,omitempty"`
	}

	PlaceOCOData struct {
		PlaceOrderData
		Other PlaceOCOOther `json:"other"`
	}
)
