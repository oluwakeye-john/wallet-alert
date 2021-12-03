package validators

import (
	"net/mail"
	"regexp"
)

func VerifyETHAddress(s string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(s)
}

func VerifyBTCAddress(s string) bool {
	re := regexp.MustCompile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$")
	return re.MatchString(s)
}

func VerifyLTCAddress(s string) bool {
	return VerifyBTCAddress(s)
}

func VerifyDOGEAddress(s string) bool {
	return VerifyBTCAddress(s)
}

func VerifyDASHAddress(s string) bool {
	return VerifyBTCAddress(s)
}

func IsEmailValid(s string) bool {
	if s == "" {
		return false
	}
	_, error := mail.ParseAddress(s)
	return error == nil
}
