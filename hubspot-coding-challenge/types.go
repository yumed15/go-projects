package main

type Partners struct {
	Partner []Partner `json:"partners"`
}

type Partner struct {
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Email          string   `json:"email"`
	Country        string   `json:"country"`
	AvailableDates []string `json:"availableDates"`
}

type BestDates struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	AttendeeCount int      `json:"attendeeCount"`
	Attendees     []string `json:"attendees"`
	Name          string   `json:"name"`
	StartDate     string   `json:"startDate"`
}

type Dates struct {
	dates        map[string][]string
	bestDay      string
	countInvites int
}
