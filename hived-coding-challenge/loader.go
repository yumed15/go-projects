package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func loadData() (RouteData, VehicleData, error) {
	routesData, err := ioutil.ReadFile("data/routes.json")
	if err != nil {
		return RouteData{}, VehicleData{}, errors.New("error reading routes file")
	}

	var routes RouteData
	err = json.Unmarshal(routesData, &routes)
	if err != nil {
		return RouteData{}, VehicleData{}, errors.New("error parsing routes file")
	}

	vehiclesData, err := ioutil.ReadFile("data/vehicles.json")
	if err != nil {
		return RouteData{}, VehicleData{}, errors.New("error reading routes file")
	}

	var vehicles VehicleData
	err = json.Unmarshal(vehiclesData, &vehicles)
	if err != nil {
		return RouteData{}, VehicleData{}, errors.New("error parsing vehicles file")
	}

	return routes, vehicles, nil
}
