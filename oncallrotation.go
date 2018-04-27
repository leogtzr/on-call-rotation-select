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
	name     string
	location OnCallerLocation
}

func (onCallPerson OnCallPerson) String() string {
	var buffer bytes.Buffer

	buffer.WriteByte('"')
	buffer.WriteString(onCallPerson.name)
	buffer.WriteByte('"')
	buffer.WriteString(" ~> ")

	switch onCallPerson.location {
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

func main() {

	dt := currentDate()
	fmt.Println(dt)

	dt = dt.AddDate(0, 0, 7)
	fmt.Println(dt)

	// Some code ...
	o := OnCallPerson{"Leo", MEX}
	fmt.Println(o)

}
