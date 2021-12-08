package currencies

import (
	"testing"
)

type Test struct {
	code  string
	valid bool
}

var test_cases = []Test{
	{
		code:  "BCY",
		valid: true,
	},
	{
		code:  "BTC",
		valid: true,
	},
	{
		code:  "OMOO",
		valid: false,
	},
	{
		code:  "DOGE",
		valid: true,
	},
	{
		code:  "ECHOKE",
		valid: false,
	},
	{
		code:  "RES",
		valid: false,
	},
}

func TestCurrencies(t *testing.T) {
	for _, i := range test_cases {
		_, err := GetCurrencyFromCode(i.code)

		if err != nil {
			if i.valid != false {
				t.Errorf("Test %s", i.code)
				t.Errorf("Should be %t", i.valid)
			}
		} else {
			if i.valid != true {
				t.Errorf("Test %s", i.code)
				t.Errorf("Should be %t", i.valid)
			}
		}
	}
}
