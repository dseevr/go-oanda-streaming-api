# Go Oanda Streaming API

WIP... but this will be an easy Go client for accessing Oanda's streaming API

## Tick formats

Heartbeats (come every 5 seconds or so):

```
{
	"time": "2016-12-20T05:55:46.064294036Z",
	"type": "HEARTBEAT"
}
```

Pricing update:

```
{
	"asks": [
		{
			"liquidity": 10000000,
			"price": "117.680"
		},
		{
			"liquidity": 10000000,
			"price": "117.682"
		}
	],
	"bids": [
		{
			"liquidity": 10000000,
			"price": "117.665"
		},
		{
			"liquidity": 10000000,
			"price": "117.663"
		}
	],
	"closeoutAsk": "117.684",
	"closeoutBid": "117.661",
	"instrument": "USD_JPY",
	"status": "tradeable",
	"time": "2016-12-20T05:55:35.676011610Z",
	"type": "PRICE"
}
```
