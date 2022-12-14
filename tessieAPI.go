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
	cars, err := getCars(token, baseURL)

	if err != nil {
		return api, err
	}

	api.Cars = cars

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
func getCars(token string, baseURL string) ([]Car, error) {

	cars := []Car{}
	//vehicles := AllVehicles{}

	url := baseURL + "/vehicles?only_active=false"

	vehicles, err := doAPICall[AllVehicles](url, token)

	if err != nil {
		return cars, err
	}

	for _, vehicle := range vehicles.Results {
		car := Car{}

		car.DisplayName = vehicle.LastState.DisplayName
		car.VIN = vehicle.Vin
		car.VehicleID = vehicle.LastState.VehicleID

		cars = append(cars, car)
	}

	return cars, nil
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

func (a API) GetStatus() (Status, error) {

	url := a.APIRoot + "/" + a.ActiveVIN + "/status"

	status, err := doAPICall[Status](url, a.Token)

	if err != nil {
		// return empty state
		return status, err
	}

	return status, nil

}

func (a API) GetLocation() (Location, error) {

	url := a.APIRoot + "/" + a.ActiveVIN + "/location"

	location, err := doAPICall[Location](url, a.Token)

	if err != nil {
		// return empty state
		return location, err
	}

	return location, nil

}

func (a API) GetTires() (Tires, error) {

	url := a.APIRoot + "/" + a.ActiveVIN + "/tire_pressure"

	tires, err := doAPICall[Tires](url, a.Token)

	if err != nil {
		// return empty state
		return tires, err
	}

	return tires, nil

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
