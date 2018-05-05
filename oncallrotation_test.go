package main

import (
	"testing"
	"time"
)

func TestIsDateEqual(t *testing.T) {
	a := time.Date(time.Now().Year(), time.February, 5, 0, 0, 0, 0, time.UTC)
	b := time.Date(time.Now().Year(), time.February, 5, 0, 0, 0, 0, time.UTC)

	if !isDateEqual(a, b) {
		t.Error("Dates should be equal excluding time")
	}
}

func TestTeamBasedOnLocation(t *testing.T) {
	team := teamBasedOnLocation(nonRandomizedTeam())
	if team == nil {
		t.Errorf("team is null")
	}
}

func TestShift(t *testing.T) {
	shift := Shift()
	if len(shift) != WeeksPerYear {
		t.Errorf("Expected: %d, got: %d", WeeksPerYear, len(shift))
	}
}

func TestDateIsWithinHoliday(t *testing.T) {
	holidays := []Holiday{
		Holiday{time.Date(time.Now().Year(), time.January, 11, 0, 0, 0, 0, time.UTC), "An amazing holiday."},
	}

	dt := time.Date(time.Now().Year(), time.January, 8, 0, 0, 0, 0, time.UTC)

	if isHoliday, _ := IsHolidayWithinShiftEstrict(holidays, dt); !isHoliday {
		t.Error("Error, there is a holiday during the shift ... ")
	}
}

func TestTruncateDateToWeekStart(t *testing.T) {
	mxHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())

	for _, h := range mxHolidays {
		truncated := truncateDateToStartingWeek(h.Date)
		if int(time.Monday) != int(truncated.Weekday()) {
			t.Error("It should be Monday")
		}
	}

	usaHolidays := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())
	for _, h := range usaHolidays {
		truncated := truncateDateToStartingWeek(h.Date)
		if int(time.Monday) != int(truncated.Weekday()) {
			t.Error("It should be Monday")
		}
	}

}
