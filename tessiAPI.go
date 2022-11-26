package tessieAPI

import (
	"fmt"
	"io/ioutil"
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

	return cars
}

func (a API) GetState() CurrentState {

	currentState := CurrentState{}

	url := a.APIRoot + "/" + a.ActiveVIN + "/state"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer <TOKEN>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return currentState

}
