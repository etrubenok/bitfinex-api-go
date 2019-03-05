package tests

import (
	"os"

	bitfinex "github.com/etrubenok/bitfinex-api-go/v1"
)

var (
	client *bitfinex.Client
)

func init() {
	key := os.Getenv("BFX_API_KEY")
	secret := os.Getenv("BFX_API_SECRET")
	client = bitfinex.NewClient().Auth(key, secret)
}
