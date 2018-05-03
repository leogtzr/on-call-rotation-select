package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// OnCallerLocation ...
type OnCallerLocation int

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

// OnCallPerson ...
type OnCallPerson struct {
	Name     string
	Location OnCallerLocation
}

// Team represent the list team members.
type Team []OnCallPerson

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

// Rotation ...
type Rotation struct {
	Date time.Time
	OnCallPerson
}

func (rotation *Rotation) String() string {
	return fmt.Sprintf("[From: %s, To: %s] - [%s]",
		rotation.Date.Format("2006-01-02"),
		rotation.Date.AddDate(0, 0, AWeek).Format("2006-01-02"), rotation.OnCallPerson.String())
}

func adjustTeamWeeks(team Team) Team {
	adjustedTeam := make([]OnCallPerson, 0)
	i := 0

	for len(adjustedTeam) <= WeeksPerYear {
		if i < len(team)-1 {
			adjustedTeam = append(adjustedTeam, team[i])
		} else {
			i = 0
		}
		i++
	}

	return adjustedTeam
}

func maxNumberOfRotations(weeksPerYear int, team Team) float64 {
	maxNumOfRotations, _ :=
		strconv.ParseFloat(fmt.Sprintf("%.0f", float64(weeksPerYear)/float64(len(team))), 64)
	return maxNumOfRotations
}

// Shift ...
func Shift() []Rotation {

	team := nonRandomizedTeam()
	maxNumOfRotations := maxNumberOfRotations(WeeksPerYear, team)
	teamShiftCounts := teamShiftsOccurrencesCount(team)

	fmt.Println(maxNumOfRotations)
	initialShiftDate := initialRotationDate()

	mxHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
	usaHolidays := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())

	shift := make([]Rotation, 0)
	for len(shift) < WeeksPerYear {
		isHolidayMX, holidayMX := IsHolidayWithinShiftEstrict(mxHolidays, initialShiftDate)
		isHolidayUSA, holidayUSA := IsHolidayWithinShiftEstrict(usaHolidays, initialShiftDate)

		if isHolidayMX && !isHolidayUSA {
			fmt.Println("There is a collision with: ", holidayMX, " but USA is free ... ")
			t := getRandomTeamWithLocation(teamShiftCounts, USA)
			fmt.Println(t, " has been selected ... ")
		}

		if isHolidayUSA && !isHolidayMX {
			fmt.Println("There is a collision with: ", holidayUSA, " but MX is free ... ")
			t := getRandomTeamWithLocation(teamShiftCounts, MEX)
			fmt.Println(t, " has been selected ... ")
		}

		if !isHolidayMX && !isHolidayUSA {
			fmt.Println("There is no collision ... ", initialShiftDate)
		}

		shift = append(shift, Rotation{})
		initialShiftDate = initialShiftDate.AddDate(0, 0, AWeek)
	}

	return shift
}

func onCallShift() []Rotation {

	rand.Seed(time.Now().Unix())

	team := nonRandomizedTeam()
	team = adjustTeamWeeks(team)
	fmt.Println("~~ ", len(team))
	team = shuffleTeam(team)

	shift := make([]Rotation, 0)
	initialShiftDate := initialRotationDate()
	mexHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
	usaHolidays := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())
	var holidays []Holiday
	nextAvailableIndex := -1

	for i := 0; i < len(team); i++ {
		t := team[i]

		if t.Location == MEX {
			holidays = mexHolidays
		} else {
			holidays = usaHolidays
		}

		if is, holiday := IsHolidayWithinShiftEstrict(holidays, initialShiftDate); is {
			fmt.Println("Collision with: ", t, ", date: ", initialShiftDate, ", holiday: ", holiday)
			if t.Location == MEX {
				nextAvailableIndex = findNextAvailableIndex(team, i, USA)
			} else {
				nextAvailableIndex = findNextAvailableIndex(team, i, MEX)
			}

			fmt.Println("Next available index is: ", nextAvailableIndex, ", which is: ", team[nextAvailableIndex], " current index is: ", i)
			if nextAvailableIndex != i {
				team[nextAvailableIndex], team[i] = team[i], team[nextAvailableIndex]
				shift = append(shift, Rotation{Date: initialShiftDate, OnCallPerson: team[i]})
			} else {
				fmt.Println("Shit ... ")
			}
		} else {
			shift = append(shift, Rotation{Date: initialShiftDate, OnCallPerson: t})
		}

		initialShiftDate = initialShiftDate.AddDate(0, 0, AWeek)
	}

	return shift

}

func findNextAvailableIndex(team Team, currentIndex int, location OnCallerLocation) int {
	for i := currentIndex + 1; i < len(team)-1; i++ {
		if team[i].Location == location {
			return i
		}
	}
	return currentIndex
}

// IsHolidayWithinShiftEstrict ...
func IsHolidayWithinShiftEstrict(holidays []Holiday, shift time.Time) (bool, *Holiday) {
	startingShift := truncateDateToStartingWeek(shift)
	endingShift := startingShift.AddDate(0, 0, AWeek)

	for _, holiday := range holidays {
		if (holiday.Date.After(startingShift) || isDateEqual(holiday.Date, startingShift)) &&
			(holiday.Date.Before(endingShift) /*|| isDateEqual(holiday.Date, endingShift)*/) {
			return true, &holiday
		}
	}

	return false, nil
}

func isDateEqual(a, b time.Time) bool {
	return (a.Year() == b.Year()) &&
		(a.Month() == b.Month()) &&
		(a.Day() == b.Day())
}

// Holiday ...
type Holiday struct {
	Date    time.Time
	Holiday string
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
		Holiday{time.Date(0, time.November, 19, 0, 0, 0, 0, time.UTC), "Revolución Mexicana"},
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
