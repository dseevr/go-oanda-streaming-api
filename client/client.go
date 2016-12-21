package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func (t *Tick) IsHeartbeat() bool {
	return "HEARTBEAT" == t.Type
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
	return fmt.Sprintf(baseUrl, c.account, c.currencies)
}

func (c *Client) Run(f func(*Tick)) {
	req, err := http.NewRequest("GET", c.url(), nil)
	if err != nil {
		log.Fatalln("http.NewRequest:", err)
		return
	}

	// set our bearer token
	req.Header.Set("Authorization", "Bearer "+c.token)

	// just use the DefaultClient, no need to be fancy here
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("http.Get:", err)
		return
	}

	tick := &Tick{}

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			// technically, we should never get io.EOF here
			log.Fatalln("reader.ReadBytes:", err)
			return
		}

		if err := json.Unmarshal(line, tick); err != nil {
			log.Fatalln("json.Unmarshal:", err)
			return
		}

		// skip the heartbeat which is sent every 5 seconds
		if tick.IsHeartbeat() {
			continue
		}

		f(tick)
	}
}
