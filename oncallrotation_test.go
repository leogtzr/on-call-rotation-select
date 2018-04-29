package main

import (
	"fmt"
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

func TestDateIsWithinHoliday(t *testing.T) {
	mexHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())

	dt := initialRotationDate()
	for i := 0; i < 50; i++ {
		if is, holiday := IsHolidayWithinShiftEstrict(mexHolidays, dt); is {
			fmt.Printf("Holiday [%s] is within date: [%s]\n\n", holiday, dt)
		}
		dt = dt.AddDate(0, 0, AWeek)
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
