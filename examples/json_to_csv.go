package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dseevr/go-oanda-streaming-api/client"
)

func main() {
	fmt.Fprint(os.Stdout, "unix_timestamp,nanoseconds,symbol,best_bid,best_ask\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tick := &client.Tick{}

		err := json.Unmarshal(scanner.Bytes(), tick)
		if err != nil {
			log.Fatalln(err)
		}

		// don't assume heartbeats were removed
		if tick.IsHeartbeat() {
			continue
		}

		var pipFormat string

		// output the correct number of pips
		if tick.IsJapanese() {
			pipFormat = "%.3f"
		} else {
			pipFormat = "%.5f"
		}

		format := "%d,%d,%s," + pipFormat + "," + pipFormat + "\n"

		fmt.Fprintf(os.Stdout, format,
			tick.UnixTimestamp(),
			tick.Nanoseconds(),
			tick.Symbol(),
			tick.BestBid(),
			tick.BestAsk(),
		)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
