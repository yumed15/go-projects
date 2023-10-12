# Optimising electricity consumption

Welcome to the Delivery Emissions live coding exercise for HIVED, the eco-friendly delivery company :truck::recycle:

At HIVED, we will never use fossil fuels, so all our vans are electric. :zap:
But different sizes of vans will vary in their electricity consumption, and so it's better for the planet (and for our wallet! :money_with_wings:) to carefully select which vans are used for which delivery routes.

## Exercise

We'd like you to write a program that assigns vehicles to routes, in a way that will minimise electricity consumption (in kWh).

## Instructions

* :speech_balloon: See this exercise as an interactive session, ask us questions as you would if we were working together
* :ok_hand: Aim to write code in the way you would every day - **you will not be penalised for not completing the exercise**


### Input

1. `routes.json` This contains an array of different routes. Each route will have an ID and a list of stops. Each stop will have a distance from the previous stop in kilometers.

2. `vehicles.json` This contains an array of vehicles. Each vehicle will have an ID, a maximum range in kilometers, and its electricity consumption in kWh/100km.

### Output

As an output, we'd like to see:

* The list of optimal vehicle-route pairs
* The total kWh required to complete all routes one after the other using the least efficient vehicle
* The total kWh required to complete all routes in parallel using the best vehicle-route pairing