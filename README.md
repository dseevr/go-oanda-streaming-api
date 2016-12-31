# Go Oanda Streaming API

This is a Go client library for Oanda's v20 streaming API.  Give it your account number, authorization token, some currencies you want live updates for, and a function to run when each tick comes in.

## Usage

```go
c := client.New(accountID, authToken, currenciesToTrack)
c.Run(func(t *client.Tick) {
	// this function fires every time a tick is received
})

```

For actual uses, see the `examples` directory.

## Examples

`printer/main.go` just prints the `client.Tick` object generated when each tick is received.

`json_to_csv/main.go` converts the JSON tick format into a CSV of `unix_timestamp,nanoseconds,symbol,best_bid,best_ask` so you can import it into other programs/languages (R, Excel, etc.).  Example invocation:

```sh
go get github.com/dseevr/go-oanda-streaming-api/examples/json_to_csv
json_to_csv <tick_data.txt >output.csv
```

`stats/main.go` reads a JSON dump from the streaming API and outputs some stats about the pairs you were tracking (currently just displays them by tick count).

## Setup

You will need to follow the instructions here to get an authorization token: http://developer.oanda.com/rest-live-v20/introduction/

Your account ID looks like `xxx-xxx-xxxxxx-xxx`.

The currency pairs you are interested in should be separated by underscores like `EUR_USD`.  If you want more than one, you'll have to separate them with a URL-encoded comma like `EUR_USD%2CEUR_JPY`.

The API being used is the pricing stream endpoint documented here: http://developer.oanda.com/rest-live-v20/pricing-ep/

Oanda will throttle or ban your IP if you connect to their servers too frequently.  Please read their documentation to see their limits and recommendations.

## Tick formats

Oanda's streaming API sends newline-separated JSON records for each tick.  There are two types of ticks.

### Heartbeats

These arrive about every 5 seconds:

```
{
	"time": "2016-12-20T05:55:46.064294036Z",
	"type": "HEARTBEAT"
}
```

### Price Updates

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

## License

BSD

## Disclaimer

No guarantees are provided or implied here.  This code is not guaranteed to even work properly.  Trusting it with something as serious as trading forex without subjecting it to a thorough review would be extremely foolish.  You will probably lose all of your money anyways, but that is not even remotely my fault whether you use my code or not.
