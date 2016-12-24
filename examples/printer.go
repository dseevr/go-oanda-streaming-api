package main

import (
	"fmt"
	"os"

	"github.com/dseevr/go-oanda-streaming-api/client"
)

func main() {
	account := os.Getenv("OANDA_ACCOUNT")
	token := os.Getenv("OANDA_TOKEN")
	currencies := os.Getenv("OANDA_CURRENCIES")

	c := client.New(account, token, currencies)
	c.Run(func(t *client.Tick) {
		fmt.Printf("%#v\n", t)
	})
}
