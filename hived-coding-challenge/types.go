package main

type RouteData struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	RouteID       string `json:"route_id"`
	Stops         []Stop `json:"stops"`
	TotalDistance float64
}

type Stop struct {
	StopID     string  `json:"stop_id"`
	DistanceKM float64 `json:"distance_km"`
}

type VehicleData struct {
	Vehicles []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	ID          string `json:"id"`
	MaxRangeKm  int64  `json:"max_range_km"`
	KWhPer100km int64  `json:"kwh_per_100_km"`
	InUse       bool
}

type VehicleRoutePair struct {
	RouteID   string
	VehicleID string
}
