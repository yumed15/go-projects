package main

import (
	"fmt"
	"time"
)

func process() {
	api := NewGatewayClient(url1, url2, apiKey)

	res, err := api.getData()
	if err != nil {
		fmt.Printf("error getting data %s", err)
	}

	availableDates, err := findBestAvailableDates(res)
	if err != nil {
		fmt.Printf("error finding available dates %s", err)
	}

	var bestDates BestDates
	for k, v := range availableDates {
		var country Country
		country.Name = k
		country.AttendeeCount = v.countInvites
		country.StartDate = v.bestDay
		country.Attendees = v.dates[v.bestDay]

		bestDates.Countries = append(bestDates.Countries, country)
	}

	err = api.sendData(bestDates)
	if err != nil {
		fmt.Printf("error sending data %s", err)
	}
}

func findBestAvailableDates(partners Partners) (map[string]Dates, error) {

	res := map[string]Dates{}

	for _, partner := range partners.Partner {
		dates, err := getDates(partner)
		if err != nil {
			return nil, err
		}

		countryDates, ok := res[partner.Country]
		if !ok {
			countryDates = Dates{dates: make(map[string][]string), bestDay: "", countInvites: 0}
		}

		for _, date := range dates {
			_, ok := countryDates.dates[date]
			if !ok {
				countryDates.dates[date] = []string{partner.Email}
			} else {
				countryDates.dates[date] = append(countryDates.dates[date], partner.Email)
			}
			if countryDates.countInvites < len(countryDates.dates[date]) {
				countryDates.countInvites = len(countryDates.dates[date])
				countryDates.bestDay = date
			}
		}
		res[partner.Country] = countryDates
	}

	return res, nil
}

func getDates(partner Partner) ([]string, error) {
	var res []string

	for i := 0; i < len(partner.AvailableDates)-1; i++ {
		ok, err := areDatesConsecutive(partner.AvailableDates[i], partner.AvailableDates[i+1])
		if err != nil {
			return []string{}, err
		}

		if ok {
			res = append(res, partner.AvailableDates[i])
		}
	}

	return res, nil
}

func areDatesConsecutive(dateStr1, dateStr2 string) (bool, error) {
	layout := "2006-01-02"
	parsedDate1, err := time.Parse(layout, dateStr1)
	if err != nil {
		return false, err
	}

	parsedDate2, err := time.Parse(layout, dateStr2)
	if err != nil {
		return false, err
	}

	duration := parsedDate2.Sub(parsedDate1)

	if duration == 24*time.Hour {
		return true, nil
	}

	return false, nil
}
