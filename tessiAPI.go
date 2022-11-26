package tessieAPI

import (
	"encoding/json"
	"io"
	"net/http"

	"errors"
)

func NewAPI(token string, baseURL string) (API, error) {
	api := API{}

	api.Token = token
	api.APIRoot = baseURL
	api.Cars = getCars(token, baseURL)

	// If there are no active cars on this account then raise an error
	if len(api.Cars) == 0 {
		return API{}, errors.New("there are no active cars on this account")
	}

	// Default to the first car in the cars list for the ActiveVIN.
	// Most API calls require a VIN
	api.ActiveVIN = api.Cars[0].VIN

	return api, nil
}

// This function grabs all the VINS assocaited with the account
func getCars(token string, baseURL string) []Car {

	cars := []Car{}
	//vehicles := AllVehicles{}

	url := baseURL + "/vehicles?only_active=false"

	vehicles, err := doAPICall[AllVehicles](url, token)

	if err != nil {
		return cars
	}

	for _, vehicle := range vehicles.Results {
		car := Car{}

		car.DisplayName = vehicle.LastState.DisplayName
		car.VIN = vehicle.Vin
		car.VehicleID = vehicle.LastState.VehicleID

		cars = append(cars, car)
	}

	return cars
}

func (a API) GetState() (CurrentState, error) {

	url := a.APIRoot + "/" + a.ActiveVIN + "/state"

	currentState, err := doAPICall[CurrentState](url, a.Token)

	if err != nil {
		// return empty state
		return currentState, err
	}

	return currentState, nil

}

// Generic function that performs the API call
func doAPICall[T APITypes](apiEndpoint string, token string) (T, error) {

	retObj := new(T)

	req, _ := http.NewRequest("GET", apiEndpoint, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		return *retObj, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, retObj)

	return *retObj, nil

}
