package customerrors

import (
	"fmt"
)

func UnsupportedCurrency(code string) error {
	return fmt.Errorf("currency %s is not supported", code)
}

func InvalidAddress() error {
	return fmt.Errorf("your wallet address is invalid")
}

func InvalidEmail() error {
	return fmt.Errorf("please enter a valid email address")
}

func AlreadySubscribed() error {
	return fmt.Errorf("you are already subscribed")
}
