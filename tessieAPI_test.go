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

func TestGetState(t *testing.T) {
	TOKEN := os.Getenv("TESSIE_TOKEN")
	API, err := NewAPI(TOKEN, "https://api.tessie.com")

	if assert.Nil(t, err) {
		state, err := API.GetState()

		if assert.Nil(t, err) {
			assert.NotEmpty(t, state.VehicleID, "No vehicleID")
		}
	}
}

func TestGetLocation(t *testing.T) {
	TOKEN := os.Getenv("TESSIE_TOKEN")
	API, err := NewAPI(TOKEN, "https://api.tessie.com")

	if assert.Nil(t, err) {
		location, err := API.GetLocation()

		if assert.Nil(t, err) {
			assert.NotEmpty(t, location.Latitude, "No latitude returned")
		}
	}
}

func TestGetStatus(t *testing.T) {
	TOKEN := os.Getenv("TESSIE_TOKEN")
	API, err := NewAPI(TOKEN, "https://api.tessie.com")

	if assert.Nil(t, err) {
		status, err := API.GetStatus()

		if assert.Nil(t, err) {
			assert.NotEmpty(t, status.Status, "No status returned")
		}
	}
}

func TestGetTires(t *testing.T) {
	TOKEN := os.Getenv("TESSIE_TOKEN")
	API, err := NewAPI(TOKEN, "https://api.tessie.com")

	if assert.Nil(t, err) {
		tires, err := API.GetTires()

		if assert.Nil(t, err) {
			assert.NotEmpty(t, tires.FrontLeftStatus, "One tire status didnt return a value")
		}
	}
}
