package main

import (
	"bytes"
	"fmt"
	"time"
)

// OnCallerLocation ...
type OnCallerLocation int

const (
	// USA ...
	USA OnCallerLocation = 0
	// MEX ...
	MEX OnCallerLocation = 1
)

// OnCallPerson ...
type OnCallPerson struct {
	Name     string
	Location OnCallerLocation
}

// Holiday ...
type Holiday struct {
	Date    time.Time
	Holiday string
}

func (holiday Holiday) String() string {
	return fmt.Sprintf("%s -> %s", holiday.Date, holiday.Holiday)
}

func buildUSAHolidays() []Holiday {
	return []Holiday{
		Holiday{time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC), "New Year's Day"},
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

func currentDate() time.Time {
	h, min, s, nsec := 0, 0, 0, 0
	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		h,
		min,
		s,
		nsec,
		time.UTC,
	)
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

func main() {

	dt := currentDate()
	fmt.Println(dt)

	dt = dt.AddDate(0, 0, 7)
	fmt.Println(dt)

	// Some code ...
	o := OnCallPerson{"Leo", MEX}
	fmt.Println(o)

	initialRotationDate := initialRotationDate()
	fmt.Println(initialRotationDate)

	fmt.Println(buildUSAHolidays())

}
