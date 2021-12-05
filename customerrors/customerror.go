package customerrors

import (
	"fmt"
)

func UnsupportedCurrency(code string) error {
	return fmt.Errorf("currency %s is not supported", code)
}

func InvalidAddress() error {
	return fmt.Errorf("address is invalid")
}

func InvalidEmail() error {
	return fmt.Errorf("invalid email address")
}

func AlreadySubscribed() error {
	return fmt.Errorf("already subscribed")
}

func NotSubscribed() error {
	return fmt.Errorf("not subscribed")
}
