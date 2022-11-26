package tessieAPI

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	TOKEN := os.Getenv("TESSIE_TOKEN")

	assert.NotEqual(t, TOKEN, "", "Token should not be an empty string. Set the TESSIE_TOKEN environment variable")

	API, err := NewAPI(TOKEN, "https://api.tessie.com")

	if assert.Nil(t, err) {
		assert.Greater(t, len(API.Cars), 0, "API creation suceeded but there are no active vehicles on the account")
	}

}
