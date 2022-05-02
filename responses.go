package go_tradovate_client

type (
	Response interface{}

	GetAccessTokenResponse struct {
		Response
		AccessToken               string `json:"accessToken"`
		MDAccessToken             string `json:"mdAccessToken"`
		ExpirationTime            string `json:"expirationTime"`
		UserStatus                string `json:"userStatus"`
		UserID                    int    `json:"userId"`
		Name                      string `json:"name"`
		HasLive                   bool   `json:"hasLive"`
		OutdatedTAC               bool   `json:"outdatedTaC"`
		HasFunded                 bool   `json:"hasFunded"`
		HasMarketData             bool   `json:"hasMarketData"`
		OutdatedLiquidationPolicy bool   `json:"outdatedLiquidationPolicy"`
	}

	PlaceOrderResponse struct {
		Response
	}

	CancelOrderResponse struct {
		Response
	}

	PlaceOCOResponse struct {
		Response
	}
)
