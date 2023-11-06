Program that assigns vehicles to routes, in a way that will minimise electricity consumption (in kWh).

### Input

1. `routes.json` This contains an array of different routes. Each route will have an ID and a list of stops. Each stop will have a distance from the previous stop in kilometers.

2. `vehicles.json` This contains an array of vehicles. Each vehicle will have an ID, a maximum range in kilometers, and its electricity consumption in kWh/100km.

### Output

As an output, we'd like to see:

* The list of optimal vehicle-route pairs
* The total kWh required to complete all routes one after the other using the least efficient vehicle
* The total kWh required to complete all routes in parallel using the best vehicle-route pairing