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
	re := regexp.MustCompile("^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$")
	return re.MatchString(s)
}

func VerifyDOGEAddress(s string) bool {
	re := regexp.MustCompile("^D{1}[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}$")
	return re.MatchString(s)
}

func VerifyDASHAddress(s string) bool {
	re := regexp.MustCompile("^X[1-9A-HJ-NP-Za-km-z]{33}")
	return re.MatchString(s)
}

func VerifyTestAddress(s string) bool {
	re := regexp.MustCompile("^(B|C|D)[a-zA-HJ-NP-Z0-9]{25,39}$")
	return re.MatchString(s)
}

func IsEmailValid(s string) bool {
	if s == "" {
		return false
	}
	_, error := mail.ParseAddress(s)
	return error == nil
}
