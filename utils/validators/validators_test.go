package validators

import (
	"testing"
)

type Test struct {
	input     string
	validator func(string) bool
	want      bool
}

var test_cases = []Test{
	// BTC
	{
		input:     "1BoatSLRHtKNngkdXEeobR76b53LETtpyT", // btc
		validator: VerifyBTCAddress,
		want:      true,
	},
	{
		input:     "soo-wrong", // dummy
		validator: VerifyBTCAddress,
		want:      false,
	},

	// BCY
	{
		input:     "CBJV4XuU1HRacTfTXES1JFHcWVsJ52EeNh", // bcy
		validator: VerifyTestAddress,
		want:      true,
	},
	{
		input:     "1BoatSLRHtKNngkdXEeobR76b53LETtpyT", // btc
		validator: VerifyTestAddress,
		want:      false,
	},

	// ETH
	{
		input:     "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // eth
		validator: VerifyETHAddress,
		want:      true,
	},
	{
		input:     "CBJV4XuU1HRacTfTXES1JFHcWVsJ52EeNh", // bcy
		validator: VerifyETHAddress,
		want:      false,
	},

	// DOGE
	{
		input:     "DLCDJhnh6aGotar6b182jpzbNEyXb3C361", // eth
		validator: VerifyDOGEAddress,
		want:      true,
	},
	{
		input:     "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // bcy
		validator: VerifyDOGEAddress,
		want:      false,
	},
}

func TestValidators(t *testing.T) {
	for _, i := range test_cases {

		is_valid := i.validator(i.input)

		if is_valid != i.want {
			t.Errorf("Test %s", i.input)
			t.Errorf("Should be %t", i.want)
		}
	}
	//
}
