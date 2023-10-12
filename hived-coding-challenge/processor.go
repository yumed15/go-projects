package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func process() error {
	routes, vehicles, err := loadData()
	if err != nil {
		return err
	}

	pairs := assignVehiclesToRoutes(routes, vehicles)
	fmt.Println("Optimal Routes:")
	for _, pair := range pairs {
		fmt.Printf("%s %s\n", pair.RouteID, pair.VehicleID)
	}

	return nil
}

func assignVehiclesToRoutes(routes RouteData, vehicles VehicleData) []VehicleRoutePair {
	var routePairs []VehicleRoutePair

	for i, route := range routes.Routes {
		routes.Routes[i].TotalDistance = calculateTotalDistanceForRoute(route)
	}

	compFunc := func(a, b Route) int {
		if a.TotalDistance < b.TotalDistance {
			return 1
		}
		return -1
	}

	slices.SortFunc(routes.Routes, compFunc)

	for _, route := range routes.Routes {
		vehicle, i := findOptimalVehicle(route, vehicles.Vehicles)
		vehicles.Vehicles[i].InUse = true
		routePairs = append(routePairs, VehicleRoutePair{VehicleID: vehicle.ID, RouteID: route.RouteID})
	}

	return routePairs
}

func findOptimalVehicle(route Route, vehicles []Vehicle) (Vehicle, int) {
	var minKWH float64
	var optimalVehicle Vehicle
	var index int

	for i, vehicle := range vehicles {
		if route.TotalDistance > float64(vehicle.MaxRangeKm) || vehicle.InUse == true {
			continue
		}

		kwh := route.TotalDistance * float64(vehicle.KWhPer100km) / 100
		if kwh < minKWH || minKWH == 0 {
			minKWH = kwh
			optimalVehicle = vehicle
			index = i
		}
	}

	return optimalVehicle, index
}

func calculateTotalDistanceForRoute(route Route) float64 {
	totalDistance := 0.0

	for _, stop := range route.Stops {
		totalDistance += stop.DistanceKM
	}

	return totalDistance
}
