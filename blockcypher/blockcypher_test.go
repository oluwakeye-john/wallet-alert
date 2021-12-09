package blockcypher

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
)

func TestCreateTestAddress(t *testing.T) {
	error := godotenv.Load("../.env")
	if error != nil {
		log.Fatalf("Error loading env")
	}

	addr, err := CreateTestAddress()

	if err != nil {
		t.Fatalf("Should create a test address")
	}

	isValid := validators.VerifyTestAddress(addr.Address)

	if !isValid {
		t.Fatalf("Should be a valid test address")
	}
}

func containsHook(hook_id string, currency_code string) bool {
	hooks, error := ListAllHooks(currency_code)

	if error != nil {
		return false
	}

	found := false

	for _, i := range hooks {
		if i.ID == hook_id {
			log.Println("Found", i.ID)
			found = true
		}
	}

	return found
}

func TestHooks(t *testing.T) {
	res, err := CreateTestAddress()

	address := res.Address
	currency_code := res.CurrencyCode.String()

	if err != nil {
		t.Fatalf("Should create a test address")
	}

	// create hook
	hook, error := SetupAddressTransactionHook(address, currency_code)

	if error != nil {
		t.Fatalf("Should create an address hook")
	}

	if !containsHook(hook.ID, currency_code) {
		t.Fatalf("New hook should be part of the returned hooks")
	}

	// delete hook
	error = DeleteAddressTransactionHook(hook.ID, currency_code)
	if error != nil {
		t.Fatalf("Should delete the hook")
	}

	if containsHook(hook.ID, currency_code) {
		t.Fatalf("New hook should be deleted")
	}
}
