package client

import "fmt"

const (
	baseUrl = "https://stream-fxtrade.oanda.com/v3/accounts/%s/pricing/stream?instruments=%s"
)

// {
// 	"time": "2016-12-20T05:55:46.064294036Z",
// 	"type": "HEARTBEAT"
// }

// {
// 	"asks": [
// 		{
// 			"liquidity": 10000000,
// 			"price": "117.680"
// 		},
// 		{
// 			"liquidity": 10000000,
// 			"price": "117.682"
// 		}
// 	],
// 	"bids": [
// 		{
// 			"liquidity": 10000000,
// 			"price": "117.665"
// 		},
// 		{
// 			"liquidity": 10000000,
// 			"price": "117.663"
// 		}
// 	],
// 	"closeoutAsk": "117.684",
// 	"closeoutBid": "117.661",
// 	"instrument": "USD_JPY",
// 	"status": "tradeable",
// 	"time": "2016-12-20T05:55:35.676011610Z",
// 	"type": "PRICE"
// }

type Tick struct {
	Asks        []Quote `json:"asks"`
	Bids        []Quote `json:"bids"`
	CloseoutAsk string  `json:"closeoutAsk"`
	CloseoutBid string  `json:"closeoutBid"`
	Instrument  string  `json:"instrument"`
	Status      string  `json:"status"`
	Time        string  `json:"time,omitempty"`
	Type        string  `json:"type"`
}

type Quote struct {
	Liquidity int64  `json:"liquidity"`
	Price     string `json:"price"`
}

type Client struct {
	account    string
	token      string
	currencies string
}

func New(account, token, currencies string) *Client {
	return &Client{
		account:    account,
		token:      token,
		currencies: currencies,
	}
}

func (c *Client) url() string {
	return fmt.Sprintf(baseUrl, account, currencies)
}

func (c *Client) Run(f func(*Tick)) {
	// TODO: connect to streaming server using token
	// TODO: stream ticks
}
