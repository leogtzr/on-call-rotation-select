package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRotationGeneration(t *testing.T) {
	shift := buildOnCallShift()
	if shift == nil {
		t.Error("shift is empty ... ")
	}
}

func TestDateIsWithinHoliday(t *testing.T) {
	mexHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())

	dt := initialRotationDate()
	for i := 0; i < 50; i++ {
		if is, holiday := IsHolidayWithinShift(mexHolidays, dt); is {
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
