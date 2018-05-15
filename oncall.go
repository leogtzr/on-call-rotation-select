package oncall

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
	USA onCallerLocation = 0
	// MEX ...
	MEX onCallerLocation = 1

	aWeek        = 7
	weeksPerYear = 52
)

// OnCallerLocation ...
type onCallerLocation int

// Rotation ...
type Rotation struct {
	Date time.Time
	Person
}

// Person ...
type Person struct {
	Name     string
	Location onCallerLocation
}

// Team represent the list team members.
type Team []Person

// Holiday ...
type Holiday struct {
	Date    time.Time
	Holiday string
}

func nonRandomizedTeam() Team {
	return []Person{
		Person{Name: "Shxin", Location: USA},
		Person{Name: "Abhi", Location: USA},
		Person{Name: "Paulina", Location: USA},
		Person{Name: "Brodr", Location: USA},
		Person{Name: "Jing", Location: USA},
		Person{Name: "Jieru", Location: USA},
		Person{Name: "Smit", Location: USA},
		Person{Name: "Him", Location: USA},
		Person{Name: "Manj", Location: USA},
		Person{Name: "Andrew", Location: USA},
		Person{Name: "Markos", Location: USA},
		Person{Name: "KritiSr", Location: USA},
		Person{Name: "AndresM", Location: MEX},
		Person{Name: "AndresD", Location: MEX},
		Person{Name: "Cizar", Location: MEX},
		Person{Name: "Hanzel", Location: MEX},
		Person{Name: "Janci", Location: MEX},
		Person{Name: "Pp", Location: MEX},
		Person{Name: "MiKaik", Location: MEX},
		Person{Name: "Alvert", Location: MEX},
		Person{Name: "Marielix", Location: MEX},
		Person{Name: "DinisR", Location: MEX},
		Person{Name: "Juancho", Location: MEX},
		Person{Name: "MiRober", Location: MEX},
		Person{Name: "MiTrivi", Location: MEX},
		Person{Name: "Javier", Location: MEX},
		Person{Name: "David", Location: MEX},
		Person{Name: "DianF", Location: MEX},
		Person{Name: "Gabo", Location: MEX},
		Person{Name: "Paco", Location: MEX},
	}
}

func teamBasedOnLocation(team Team) map[onCallerLocation]Team {

	teams := make(map[onCallerLocation]Team)
	mxTeam := make([]Person, 0)
	usaTeam := make([]Person, 0)

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

func teamShiftsOccurrencesCount(team Team) map[Person]int {
	occurrences := make(map[Person]int)
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

func getRandomTeamWithLocation(counts map[Person]int, location onCallerLocation,
) Person {
	t := getRandomTeamMember(counts)
	for t.Location != location {
		t = getRandomTeamMember(counts)
	}
	return t
}

func getRandomTeamMember(counts map[Person]int) Person {
	i := rand.Intn(len(counts))
	for k := range counts {
		if i == 0 {
			return k
		}
		i--
	}
	panic("never...")
}

func getRandomTeamMemberMax(counts map[Person]int, max int) Person {
	found := false
	var t Person
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
		rotation.Date.AddDate(0, 0, aWeek).Format("2006-01-02"), rotation.Person.String())
}

func maxNumberOfRotations(weeksPerYear int, team Team) float64 {
	maxNumOfRotations, _ :=
		strconv.ParseFloat(fmt.Sprintf("%.0f", float64(weeksPerYear)/float64(len(team))), 64)
	return maxNumOfRotations
}

func smallest(counts map[Person]int) (int, Person) {
	small := math.MaxInt64
	person := Person{}
	for k, v := range counts {
		if v < small {
			person, small = k, v
		}
	}
	return small, person
}

func smallestWithLocation(counts map[Person]int, location onCallerLocation) (int, Person) {
	small := math.MaxInt64
	person := Person{}
	for k, v := range counts {
		if v < small && k.Location == location {
			person, small = k, v
		}
	}
	return small, person
}

func everybodyHadSameShifts(counts map[Person]int, smallest int) (bool, Person) {
	for k, v := range counts {
		if v != smallest {
			return false, k
		}
	}
	return true, Person{}
}

func assignTeamMember(
	counts map[Person]int,
	maxNumOfRotations int,
	location onCallerLocation,
	shift []Rotation,
	shiftDate time.Time,
) Person {
	_, onCallPerson := smallestWithLocation(counts, location)
	counts[onCallPerson]++

	return onCallPerson
}

// Shift ...
func Shift() []Rotation {

	team := nonRandomizedTeam()
	maxNumOfRotations := maxNumberOfRotations(weeksPerYear, team)
	teamShiftCounts := teamShiftsOccurrencesCount(team)

	shiftDate := initialRotationDate()

	mxHolidays := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
	usaHolidays := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())
	var t Person

	shift := make([]Rotation, 0)
	for len(shift) < weeksPerYear {

		isHolidayMX, _ := IsHolidayWithinShiftEstrict(mxHolidays, shiftDate)
		isHolidayUSA, _ := IsHolidayWithinShiftEstrict(usaHolidays, shiftDate)

		if isHolidayMX && isHolidayUSA {
			//fmt.Printf("Collision in both sides: %v, holidayMX: %v\n", holidayUSA, holidayMX)
			_, t = smallest(teamShiftCounts)
			// fmt.Printf("Chosen: %v -> {%v} and {%v}\n", t.String(), holidayUSA, holidayMX)
			teamShiftCounts[t]++
		} else if isHolidayMX && !isHolidayUSA {
			//fmt.Printf("There is a collision with [%v], but USA is free\n", holidayMX)
			t = assignTeamMember(teamShiftCounts, int(maxNumOfRotations), USA, shift, shiftDate)
		} else if !isHolidayMX && isHolidayUSA {
			//fmt.Printf("There is a collision with [%v], but MX is free\n", holidayUSA)
			t = assignTeamMember(teamShiftCounts, int(maxNumOfRotations), MEX, shift, shiftDate)
		} else if !isHolidayMX && !isHolidayUSA {
			t = getRandomTeamMember(teamShiftCounts)
			//fmt.Println("There is no collision, we have chosen", t, " | date: ", initialShiftDate)
			teamShiftCounts[t]++
		}

		shift = append(shift, Rotation{Date: shiftDate, Person: t})
		shiftDate = shiftDate.AddDate(0, 0, aWeek)
	}

	return shift
}

// IsHolidayWithinShiftEstrict ...
func IsHolidayWithinShiftEstrict(holidays []Holiday, shift time.Time) (bool, *Holiday) {
	startingShift := truncateDateToStartingWeek(shift)
	endingShift := startingShift.AddDate(0, 0, aWeek)

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

func (h Holiday) String() string {
	return fmt.Sprintf("%s -> %s", h.Date, h.Holiday)
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

func (p Person) String() string {
	var buffer bytes.Buffer

	buffer.WriteByte('"')
	buffer.WriteString(p.Name)
	buffer.WriteByte('"')
	buffer.WriteString(" ~> ")

	switch p.Location {
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
