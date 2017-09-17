package main

import (
	"log"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

func main() {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
	}
	client.APIVersion = "2015-11-06"

	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            "OMISE_GO Test Card",
		Number:          "4242424242424242",
		ExpirationMonth: 12,
		ExpirationYear:  2018,
	}
	if e := client.Do(token, createToken); e != nil {
		log.Fatal(e)
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   100000, // ฿ 1,000.00
		Currency: "thb",
		Card:     token.ID,
	}
	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}
