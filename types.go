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

// Team represent the list team members.
type Team []OnCallPerson

func nonRandomizedTeam() Team {
	return []OnCallPerson{
		OnCallPerson{Name: "Shxin", Location: USA},
		OnCallPerson{Name: "Abhi", Location: USA},
		OnCallPerson{Name: "Paul", Location: USA},
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

// Rotation ...
type Rotation struct {
	Date time.Time
	OnCallPerson
}

// Holiday ...
type Holiday struct {
	Date    time.Time
	Holiday string
}

func (hd Holiday) String() string {
	return fmt.Sprintf("%s -> %s", hd.Date, hd.Holiday)
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

func buildMEXHolidays() []Holiday {
	return []Holiday{
		Holiday{time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC), "Año Nuevo"},
		Holiday{time.Date(0, time.February, 5, 0, 0, 0, 0, time.UTC), "Día de la Constitución Mexicana"},
		Holiday{time.Date(0, time.March, 19, 0, 0, 0, 0, time.UTC), "Natalicio de Benito Juárez"},
		Holiday{time.Date(0, time.May, 1, 0, 0, 0, 0, time.UTC), "Día del Trabajo"},
		Holiday{time.Date(0, time.September, 16, 0, 0, 0, 0, time.UTC), "Día de la Independencia"},
		Holiday{time.Date(0, time.November, 19, 0, 0, 0, 0, time.UTC), "Revolución Mexicana"},
		Holiday{time.Date(0, time.December, 1, 0, 0, 0, 0, time.UTC), "Transmisión de Poder Ejecutivo Federal"},
		Holiday{time.Date(0, time.December, 25, 0, 0, 0, 0, time.UTC), "Día de Navidad"},
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
