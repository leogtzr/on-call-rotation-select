package oncall

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

func TestTeamBasedOnLocation(t *testing.T) {
	team := teamBasedOnLocation(nonRandomizedTeam())
	if team == nil {
		t.Errorf("team is null")
	}
}

func TestShift(t *testing.T) {
	shift := Shift()
	if len(shift) != weeksPerYear {
		t.Errorf("Expected: %d, got: %d", weeksPerYear, len(shift))
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

func Test_teamShiftsOccurrencesCount(t *testing.T) {
	occurrences := teamShiftsOccurrencesCount(nonRandomizedTeam())
	if occurrences == nil {
		t.Error("occurrences is nil ... ")
	}
	for _, v := range occurrences {
		if v != 0 {
			t.Fatal("Initial values should be zero ... ")
		}
	}
}

func Test_getRandomTeamMember(t *testing.T) {
	team := nonRandomizedTeam()
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	randomTeamMember := getRandomTeamMember(teamShiftCounts)
	fmt.Println(randomTeamMember)
}

func Test_getRandomTeamWithLocation(t *testing.T) {
	team := nonRandomizedTeam()
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	randomTeamMember := getRandomTeamWithLocation(teamShiftCounts, USA)
	if randomTeamMember.Location != USA {
		t.Error("Error, expecting USA team member.")
	}
}

func Test_smallest(t *testing.T) {
	team := Team{
		Person{Name: "Shxin", Location: USA},
		Person{Name: "Abhi", Location: USA},
		Person{Name: "Paulina", Location: USA},
		Person{Name: "Brodr", Location: USA},
	}
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	teamShiftCounts[team[0]] = 3
	teamShiftCounts[team[1]] = 56
	teamShiftCounts[team[2]] = 5
	teamShiftCounts[team[3]] = 4556

	small, _ := smallest(teamShiftCounts)
	expected := 3
	if small != expected {
		t.Errorf("Got: %d, expected: %d", small, expected)
	}

}

func Test_everybodyHadSameShifts(t *testing.T) {
	team := Team{
		Person{Name: "Shxin", Location: USA},
		Person{Name: "Abhi", Location: USA},
		Person{Name: "Paulina", Location: USA},
		Person{Name: "Brodr", Location: USA},
	}
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	teamShiftCounts[team[0]] = 2
	teamShiftCounts[team[1]] = 2
	teamShiftCounts[team[2]] = 2
	teamShiftCounts[team[3]] = 2

	fmt.Println(teamShiftCounts)

	small, _ := smallest(teamShiftCounts)

	if ok, _ := everybodyHadSameShifts(teamShiftCounts, small); !ok {
		t.Error("Expecting true ... ")
	}
}
