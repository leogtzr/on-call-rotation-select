package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	// USA ...
	USA OnCallerLocation = 0
	// MEX ...
	MEX OnCallerLocation = 1
	// AWeek ...
	AWeek = 7
	// WeeksPerYear ...
	WeeksPerYear = 52
)

// OnCallerLocation ...
type OnCallerLocation int

// Rotation ...
type Rotation struct {
	Date time.Time
	OnCallPerson
}

// OnCallPerson ...
type OnCallPerson struct {
	Name     string
	Location OnCallerLocation
}

// Team represent the list team members.
type Team []OnCallPerson

// Holiday ...
type Holiday struct {
	Date    time.Time
	Holiday string
}

func nonRandomizedTeam() Team {
	return []OnCallPerson{
		OnCallPerson{Name: "Shxin", Location: USA},
		OnCallPerson{Name: "Abhi", Location: USA},
		OnCallPerson{Name: "Paulina", Location: USA},
		OnCallPerson{Name: "Brodr", Location: USA},
		OnCallPerson{Name: "Jing", Location: USA},
		OnCallPerson{Name: "Jieru", Location: USA},
		OnCallPerson{Name: "Smit", Location: USA},
		OnCallPerson{Name: "Him", Location: USA},
		OnCallPerson{Name: "Manj", Location: USA},
		OnCallPerson{Name: "Andrew", Location: USA},
		OnCallPerson{Name: "Markos", Location: USA},
		OnCallPerson{Name: "KritiSr", Location: USA},
		OnCallPerson{Name: "AndresM", Location: MEX},
		OnCallPerson{Name: "AndresD", Location: MEX},
		OnCallPerson{Name: "Cizar", Location: MEX},
		OnCallPerson{Name: "Hanzel", Location: MEX},
		OnCallPerson{Name: "Janci", Location: MEX},
		OnCallPerson{Name: "Pp", Location: MEX},
		OnCallPerson{Name: "MiKaik", Location: MEX},
		OnCallPerson{Name: "Alvert", Location: MEX},
		OnCallPerson{Name: "Marielix", Location: MEX},
		OnCallPerson{Name: "DinisR", Location: MEX},
		OnCallPerson{Name: "Juancho", Location: MEX},
		OnCallPerson{Name: "MiRober", Location: MEX},
		OnCallPerson{Name: "MiTrivi", Location: MEX},
		OnCallPerson{Name: "Javier", Location: MEX},
		OnCallPerson{Name: "David", Location: MEX},
		OnCallPerson{Name: "DianF", Location: MEX},
		OnCallPerson{Name: "Gabo", Location: MEX},
		OnCallPerson{Name: "Paco", Location: MEX},
	}
}

func teamBasedOnLocation(team Team) map[OnCallerLocation]Team {

	teams := make(map[OnCallerLocation]Team)
	mxTeam := make([]OnCallPerson, 0)
	usaTeam := make([]OnCallPerson, 0)

	for _, t := range team {
		if t.Location == MEX {
			mxTeam = append(mxTeam, t)
		} else {
			usaTeam = append(usaTeam, t)
		}
	}
	teams[MEX] = mxTeam
	teams[USA] = usaTeam

	return teams

}

func teamShiftsOccurrencesCount(team Team) map[OnCallPerson]int {
	occurrences := make(map[OnCallPerson]int)
	for _, t := range team {
		occurrences[t] = 0
	}
	return occurrences
}

func shuffleTeam(team Team) Team {
	for i := range team {
		j := rand.Intn(i + 1)
		team[i], team[j] = team[j], team[i]
	}
	return team
}

func getRandomTeamWithLocation(counts map[OnCallPerson]int, location OnCallerLocation,
) OnCallPerson {
	t := getRandomTeamMember(counts)
	for t.Location != location {
		t = getRandomTeamMember(counts)
	}
	return t
}

func getRandomTeamMember(counts map[OnCallPerson]int) OnCallPerson {
	i := rand.Intn(len(counts))
	for k := range counts {
		if i == 0 {
			return k
		}
		i--
	}
	panic("never...")
}

func getRandomTeamMemberMax(counts map[OnCallPerson]int, max int) OnCallPerson {
	found := false
	var t OnCallPerson
	for !found {
		t = getRandomTeamMember(counts)
		if counts[t] <= max {
			found = true
		}
	}
	return t
}

func (rotation *Rotation) String() string {
	return fmt.Sprintf("[From: %s, To: %s] - [%s]",
		rotation.Date.Format("2006-01-02"),
		rotation.Date.AddDate(0, 0, AWeek).Format("2006-01-02"), rotation.OnCallPerson.String())
}

func maxNumberOfRotations(weeksPerYear int, team Team) float64 {
	maxNumOfRotations, _ :=
		strconv.ParseFloat(fmt.Sprintf("%.0f", float64(weeksPerYear)/float64(len(team))), 64)
	return maxNumOfRotations
}

func smallest(counts map[OnCallPerson]int) (int, OnCallPerson) {
	small := math.MaxInt64
	onCallPerson := OnCallPerson{}
	for k, v := range counts {
		if v < small {
			onCallPerson, small = k, v
		}
	}
	return small, onCallPerson
}

func smallestWithLocation(counts map[OnCallPerson]int, location OnCallerLocation) (int, OnCallPerson) {
	small := math.MaxInt64
	onCallPerson := OnCallPerson{}
	for k, v := range counts {
		if v < small && k.Location == location {
			onCallPerson, small = k, v
		}
	}
	return small, onCallPerson
}

func everybodyHadSameShifts(counts map[OnCallPerson]int, smallest int) (bool, OnCallPerson) {
	for k, v := range counts {
		if v != smallest {
			return false, k
		}
	}
	return true, OnCallPerson{}
}

func assignTeamMember(
	counts map[OnCallPerson]int,
	maxNumOfRotations int,
	location OnCallerLocation,
	shift []Rotation,
	shiftDate time.Time,
) OnCallPerson {
	_, onCallPerson := smallestWithLocation(counts, location)
	counts[onCallPerson]++

	/*
			found := false
			for !found {
				ok, _ := everybodyHadSameShifts(counts, smallest)
				if !ok && onCallPerson.Location == location {
					counts[onCallPerson]++
					found = true
				} else {
					found = true
				}
		    }
	*/

	//shift = append(shift, Rotation{Date: shiftDate, OnCallPerson: onCallPerson})
	return onCallPerson
}

// Shift ...
func Shift() []Rotation {

	team := nonRandomizedTeam()
	maxNumOfRotations := maxNumberOfRotations(WeeksPerYear, team)
	teamShiftCounts := teamShiftsOccurrencesCount(team)

	initialShiftDate := initialRotationDate()

	mxHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
	usaHolidays := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())
	var t OnCallPerson

	shift := make([]Rotation, 0)
	for len(shift) < WeeksPerYear {

		isHolidayMX, holidayMX := IsHolidayWithinShiftEstrict(mxHolidays, initialShiftDate)
		isHolidayUSA, holidayUSA := IsHolidayWithinShiftEstrict(usaHolidays, initialShiftDate)

		if isHolidayMX && isHolidayUSA {
			//fmt.Printf("Collision in both sides: %v, holidayMX: %v\n", holidayUSA, holidayMX)
			_, t = smallest(teamShiftCounts)
			fmt.Printf("Chosen: %v -> {%v} and {%v}\n", t.String(), holidayUSA, holidayMX)
			teamShiftCounts[t]++
		} else if isHolidayMX && !isHolidayUSA {
			//fmt.Printf("There is a collision with [%v], but USA is free\n", holidayMX)
			t = assignTeamMember(teamShiftCounts, int(maxNumOfRotations), USA, shift, initialShiftDate)
		} else if !isHolidayMX && isHolidayUSA {
			//fmt.Printf("There is a collision with [%v], but MX is free\n", holidayUSA)
			t = assignTeamMember(teamShiftCounts, int(maxNumOfRotations), MEX, shift, initialShiftDate)
		} else if !isHolidayMX && !isHolidayUSA {
			t = getRandomTeamMember(teamShiftCounts)
			//fmt.Println("There is no collision, we have chosen", t, " | date: ", initialShiftDate)
			teamShiftCounts[t]++
		}

		shift = append(shift, Rotation{Date: initialShiftDate, OnCallPerson: t})
		initialShiftDate = initialShiftDate.AddDate(0, 0, AWeek)
	}

	return shift
}

// IsHolidayWithinShiftEstrict ...
func IsHolidayWithinShiftEstrict(holidays []Holiday, shift time.Time) (bool, *Holiday) {
	startingShift := truncateDateToStartingWeek(shift)
	endingShift := startingShift.AddDate(0, 0, AWeek)

	for _, holiday := range holidays {
		if (holiday.Date.After(startingShift) || isDateEqual(holiday.Date, startingShift)) &&
			(holiday.Date.Before(endingShift) || isDateEqual(holiday.Date, endingShift)) {
			return true, &holiday
		}
	}

	return false, nil
}

func isDateEqual(a, b time.Time) bool {
	return (a.Year() == b.Year()) && (a.Month() == b.Month()) &&
		(a.Day() == b.Day())
}

func (hd Holiday) String() string {
	return fmt.Sprintf("%s -> %s", hd.Date, hd.Holiday)
}

func truncateDateToStartingWeek(dt time.Time) time.Time {
	return dt.AddDate(0, 0, -int(dt.Weekday())+1)
}

func buildUSAHolidays() []Holiday {
	return []Holiday{
		// Holiday{time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC), "New Year's Day"},
		Holiday{time.Date(0, time.January, 15, 0, 0, 0, 0, time.UTC), "Birthday of Martin Luther King, Jr."},
		Holiday{time.Date(0, time.February, 19, 0, 0, 0, 0, time.UTC), "Washington's Birthday"},
		Holiday{time.Date(0, time.May, 28, 0, 0, 0, 0, time.UTC), "Memorial Day"},
		Holiday{time.Date(0, time.July, 4, 0, 0, 0, 0, time.UTC), "Independence Day"},
		Holiday{time.Date(0, time.September, 3, 0, 0, 0, 0, time.UTC), "Labor Day"},
		Holiday{time.Date(0, time.October, 8, 0, 0, 0, 0, time.UTC), "Columbus Day"},
		Holiday{time.Date(0, time.November, 12, 0, 0, 0, 0, time.UTC), "Veterans Day"},
		Holiday{time.Date(0, time.November, 22, 0, 0, 0, 0, time.UTC), "Thanksgiving Day"},
		Holiday{time.Date(0, time.December, 25, 0, 0, 0, 0, time.UTC), "Christmas Day"},
	}
}

func buildMEXHolidays() []Holiday {
	return []Holiday{
		// Holiday{time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC), "Año Nuevo"},
		Holiday{time.Date(0, time.February, 5, 0, 0, 0, 0, time.UTC), "Día de la Constitución Mexicana"},
		Holiday{time.Date(0, time.March, 19, 0, 0, 0, 0, time.UTC), "Natalicio de Benito Juárez"},
		Holiday{time.Date(0, time.May, 1, 0, 0, 0, 0, time.UTC), "Día del Trabajo"},
		Holiday{time.Date(0, time.September, 16, 0, 0, 0, 0, time.UTC), "Día de la Independencia"},
		Holiday{time.Date(0, time.November, 20, 0, 0, 0, 0, time.UTC), "Revolución Mexicana"},
		Holiday{time.Date(0, time.December, 1, 0, 0, 0, 0, time.UTC), "Transmisión de Poder Ejecutivo Federal"},
		Holiday{time.Date(0, time.December, 25, 0, 0, 0, 0, time.UTC), "Día de Navidad"},
	}
}

func normalizeHolidayBasedOnCurrentYear(holidays []Holiday) []Holiday {
	for i := range holidays {
		holidays[i].Date = time.Date(
			time.Now().Year(),
			holidays[i].Date.Month(),
			holidays[i].Date.Day(),
			holidays[i].Date.Hour(),
			holidays[i].Date.Minute(),
			holidays[i].Date.Second(),
			holidays[i].Date.Nanosecond(),
			time.UTC,
		)
	}
	return holidays
}

func (onCallPerson OnCallPerson) String() string {
	var buffer bytes.Buffer

	buffer.WriteByte('"')
	buffer.WriteString(onCallPerson.Name)
	buffer.WriteByte('"')
	buffer.WriteString(" ~> ")

	switch onCallPerson.Location {
	case MEX:
		buffer.WriteString("MEX")
	case USA:
		buffer.WriteString("USA")
	}

	return fmt.Sprint(buffer.String())

}

func initialRotationDate() time.Time {
	h, min, s, nsec := 0, 0, 0, 0
	return time.Date(
		time.Now().Year(),
		time.January,
		1,
		h,
		min,
		s,
		nsec,
		time.UTC,
	)
}

func initialRotationDateWithoutYear() time.Time {
	h, min, s, nsec := 0, 0, 0, 0
	return time.Date(
		-1,
		time.January,
		1,
		h,
		min,
		s,
		nsec,
		time.UTC,
	)
}
