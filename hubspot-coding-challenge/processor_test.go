package main

import (
	"testing"
)

func TestAreDatesConsecutive(t *testing.T) {
	// Test consecutive dates
	dateStr1 := "2023-09-14"
	dateStr2 := "2023-09-15"
	consecutive, err := areDatesConsecutive(dateStr1, dateStr2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if !consecutive {
		t.Errorf("Expected dates to be consecutive, but they are not")
	}

	// Test non-consecutive dates
	dateStr1 = "2023-09-14"
	dateStr2 = "2023-09-16"
	consecutive, err = areDatesConsecutive(dateStr1, dateStr2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if consecutive {
		t.Errorf("Expected dates to be non-consecutive, but they are consecutive")
	}

	// Test invalid date format
	dateStr1 = "2023-09-14"
	dateStr2 = "2023-09-xx"
	consecutive, err = areDatesConsecutive(dateStr1, dateStr2)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestAreDatesConsecutive_EdgeCases(t *testing.T) {
	// Test same dates (not consecutive)
	dateStr1 := "2023-09-14"
	dateStr2 := "2023-09-14"
	consecutive, err := areDatesConsecutive(dateStr1, dateStr2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if consecutive {
		t.Errorf("Expected dates to be non-consecutive, but they are consecutive")
	}

	// Test very large duration (not consecutive)
	dateStr1 = "0001-01-01"
	dateStr2 = "9999-12-31"
	consecutive, err = areDatesConsecutive(dateStr1, dateStr2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if consecutive {
		t.Errorf("Expected dates to be non-consecutive, but they are consecutive")
	}
}

func TestGetDates(t *testing.T) {
	// Test consecutive dates
	consecutiveDates := []string{"2023-09-14", "2023-09-15", "2023-09-16", "2023-09-17"}
	partner := Partner{AvailableDates: consecutiveDates}
	result, err := getDates(partner)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	expected := []string{"2023-09-14", "2023-09-15", "2023-09-16"}
	assertStringSlicesEqual(t, expected, result)

	// Test non-consecutive dates
	nonConsecutiveDates := []string{"2023-09-14", "2023-09-16", "2023-09-18"}
	partner = Partner{AvailableDates: nonConsecutiveDates}
	result, err = getDates(partner)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	assertStringSlicesEmpty(t, result)

	// Test error condition
	invalidDates := []string{"2023-09-14", "2023-09-xx", "2023-09-16"}
	partner = Partner{AvailableDates: invalidDates}
	result, err = getDates(partner)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func assertStringSlicesEqual(t *testing.T, expected, actual []string) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, but got %v", expected, actual)
			return
		}
	}
}

func assertStringSlicesEmpty(t *testing.T, actual []string) {
	t.Helper()
	if len(actual) != 0 {
		t.Errorf("Expected an empty slice, but got %v", actual)
	}
}
